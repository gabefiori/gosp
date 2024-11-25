package selector

// Displays a series of options for user selection.
type Selector interface {
	Run(inputChan chan string) (string, error)
}

func New(s string, opts any) Selector {
	var fzfOpts []string

	if opts != nil {
		fzfOpts = opts.([]string)
	}

	switch s {
	case "fzf":
		return NewFzf(fzfOpts)
	case "fzy":
		return NewFzy()
	default:
		return NewFzf(fzfOpts)
	}
}
