package runner

import (
	"context"
	"time"
)

// Exit constants used to know how handle the message.
// The consumer runnig is the responsible to understand this status and handle them properly.
const (
	ExitTimeout     = -1
	ExitACK         = 0
	ExitFailed      = 1
	ExitNACK        = 3
	ExitNACKRequeue = 4
	ExitRetry       = 5
)

// Runnable represent an runnable used by consumers to handle messages.
type Runnable interface {
	Process(context.Context, []byte) int
}

type Options struct {
	Path string   `mapstructure:"path"`
	Args []string `mapstructure:"args"`
}

// Config is an composition of all options and configurations used by this runnables.
type Config struct {
	Type         string        `mapstructure:"type"`
	IgnoreOutput bool          `mapstructure:"ignore-output"`
	Options      Options       `mapstructure:"options"`
	Timeout      time.Duration `mapstructure:"timeout"`
}

// New create and return a Runnable based on the config type. if the type didn't exist an error is returned.
func New(log *zap.Logger, c Config) (Runnable, error) {
	switch c.Type {
	case "command":
		return newCommand(log, c)
	}
	return nil, errors.Errorf(
		"Invalid Runner type (\"%s\") expecting (%s)",
		c.Type,
		strings.Join([]string{"command"}, ", "))
}
