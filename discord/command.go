package discord

// CommandOptionType specifies the type of the arguments used in Command.Options
type CommandOptionType int

// Constants for each slash command option type
//goland:noinspection GoUnusedConst
const (
	CommandOptionTypeSubCommand CommandOptionType = iota + 1
	CommandOptionTypeSubCommandGroup
	CommandOptionTypeString
	CommandOptionTypeInteger
	CommandOptionTypeBoolean
	CommandOptionTypeUser
	CommandOptionTypeChannel
	CommandOptionTypeRole
	CommandOptionTypeMentionable
	CommandOptionTypeNumber
)

// Command is the base "command" model that belongs to an application.
type Command struct {
	ID                Snowflake       `json:"id,omitempty"`
	ApplicationID     Snowflake       `json:"application_id,omitempty"`
	GuildID           *Snowflake      `json:"guild_id,omitempty"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Options           []CommandOption `json:"options,omitempty"`
	DefaultPermission bool            `json:"default_permission,omitempty"`
}

// CommandOption are the arguments used in Command.Options
type CommandOption struct {
	Type        CommandOptionType     `json:"type"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Required    bool                  `json:"required,omitempty"`
	Choices     []CommandOptionChoice `json:"choices,omitempty"`
	Options     []CommandOption       `json:"options,omitempty"`
}

// AddChoice adds a new choice to the CommandOption. Value can either be a string, int or float
func (o CommandOption) AddChoice(name string, value interface{}) CommandOption {
	o.Choices = append(o.Choices, CommandOptionChoice{
		Name:  name,
		Value: value,
	})
	return o
}

// AddChoices adds multiple choices to the CommandOption. Value can either be a string, int or float
func (o CommandOption) AddChoices(choices map[string]interface{}) CommandOption {
	for name, value := range choices {
		o.Choices = append(o.Choices, CommandOptionChoice{
			Name:  name,
			Value: value,
		})
	}
	return o
}

// AddOptions adds multiple choices to the CommandOption
func (o CommandOption) AddOptions(options ...CommandOption) CommandOption {
	o.Options = append(o.Options, options...)
	return o
}

// SetRequired sets if the CommandOption is required
func (o CommandOption) SetRequired(required bool) CommandOption {
	o.Required = required
	return o
}

// CommandOptionChoice contains the data for a user using your command. Value can either be a string, int or float
type CommandOptionChoice struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// CommandCreate is used to create a Command. all fields are optional
type CommandCreate struct {
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Options           []CommandOption `json:"options,omitempty"`
	DefaultPermission bool            `json:"default_permission"`
}

// CommandUpdate is used to update an existing Command. all fields are optional
type CommandUpdate struct {
	Name              *string         `json:"name,omitempty"`
	Description       *string         `json:"description,omitempty"`
	Options           []CommandOption `json:"options,omitempty"`
	DefaultPermission *bool           `json:"default_permission,omitempty"`
}
