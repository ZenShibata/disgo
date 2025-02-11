package voice

import (
	"context"
	"errors"
	"io"
	"net"
	"time"

	"github.com/disgoorg/log"
)

var (
	// SilenceAudioFrame is a 20ms opus frame of silence.
	SilenceAudioFrame = []byte{0xF8, 0xFF, 0xFE}

	// OpusFrameSize is the size of an opus frame in milliseconds.
	OpusFrameSize int64 = 20

	// OpusStreamBuffSize is the size of the buffer for receiving one opus frame.
	OpusStreamBuffSize int64 = 4000
)

type (
	// AudioSenderCreateFunc is used to create a new AudioSender sending audio to the given Conn.
	AudioSenderCreateFunc func(logger log.Logger, provider OpusFrameProvider, conn Conn) AudioSender

	// AudioSender is used to send audio to a Conn.
	AudioSender interface {
		Open()
		Close()
	}

	// OpusFrameProvider is used to provide opus frames to an AudioSender.
	OpusFrameProvider interface {
		// ProvideOpusFrame provides an opus frame to the AudioSender.
		ProvideOpusFrame() ([]byte, error)

		// Close closes the OpusFrameProvider.
		Close()
	}
)

// NewAudioSender creates a new AudioSender sending audio from the given OpusFrameProvider to the given Conn.
func NewAudioSender(logger log.Logger, opusProvider OpusFrameProvider, conn Conn) AudioSender {
	return &defaultAudioSender{
		logger:       logger,
		opusProvider: opusProvider,
		conn:         conn,
		silentFrames: 5,
	}
}

type defaultAudioSender struct {
	logger       log.Logger
	cancelFunc   context.CancelFunc
	opusProvider OpusFrameProvider
	conn         Conn

	silentFrames      int
	sentSpeakingStop  bool
	sentSpeakingStart bool
}

func (s *defaultAudioSender) Open() {
	defer s.logger.Debug("closing audio sender")
	lastFrameSent := time.Now().UnixMilli()
	ctx, cancel := context.WithCancel(context.Background())
	s.cancelFunc = cancel
	defer cancel()
loop:
	for {
		select {
		case <-ctx.Done():
			break loop

		default:
			s.send()
			sleepTime := time.Duration(OpusFrameSize - (time.Now().UnixMilli() - lastFrameSent))
			if sleepTime > 0 {
				time.Sleep(sleepTime * time.Millisecond)
			}
			if time.Now().UnixMilli() < lastFrameSent+OpusFrameSize*3 {
				lastFrameSent += OpusFrameSize
			} else {
				lastFrameSent = time.Now().UnixMilli()
			}
		}
	}
}

func (s *defaultAudioSender) send() {
	if s.opusProvider == nil {
		return
	}
	opus, err := s.opusProvider.ProvideOpusFrame()
	if err != nil && err != io.EOF {
		s.logger.Errorf("error while reading opus frame: %s", err)
		return
	}
	if len(opus) == 0 {
		if s.silentFrames > 0 {
			if _, err = s.conn.UDP().Write(SilenceAudioFrame); err != nil {
				s.handleErr(err)
			}
			s.silentFrames--
		} else if !s.sentSpeakingStop {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err = s.conn.SetSpeaking(ctx, SpeakingFlagNone); err != nil {
				s.handleErr(err)
			}
			s.sentSpeakingStop = true
			s.sentSpeakingStart = false
		}
		return
	}

	if !s.sentSpeakingStart {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = s.conn.SetSpeaking(ctx, SpeakingFlagMicrophone); err != nil {
			s.handleErr(err)
		}
		s.sentSpeakingStart = true
		s.sentSpeakingStop = false
		s.silentFrames = 5
	}

	if _, err = s.conn.UDP().Write(opus); err != nil {
		s.handleErr(err)
	}
}

func (s *defaultAudioSender) handleErr(err error) {
	if errors.Is(err, net.ErrClosed) || errors.Is(err, ErrGatewayNotConnected) {
		s.Close()
		return
	}
	s.logger.Errorf("failed to send audio: %s", err)
}

func (s *defaultAudioSender) Close() {
	s.cancelFunc()
}
