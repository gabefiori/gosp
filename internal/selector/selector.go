package selector

import "strings"

// Type represents the different types of selectors available.
type Type uint8

const (
	UnknownType Type = iota
	FzyType
	FzfType
)

func TypeFromStr(s string) Type {
	switch strings.ToLower(s) {
	case "fzy":
		return FzyType
	case "fzf":
		return FzfType
	default:
		return UnknownType
	}
}

// Displays a series of options for user selection.
type Selector interface {
	Run(inputChan chan string) (string, error)
}

// New creates a new Selector instance based on the provided selector type and options.
func New(t Type, opts any) Selector {
	var fzfOptions []string

	if opts != nil {
		fzfOptions = opts.([]string)
	}

	switch t {
	case FzfType:
		return NewFzf(fzfOptions)
	case FzyType:
		return NewFzy()
	default:
		return NewFzf(fzfOptions)
	}
}
