package selector

import (
	fzf "github.com/junegunn/fzf/src"
)

// Fzf selector for command-line fuzzy finding.
//
// For more information, see:
// https://junegunn.github.io/fzf/tips/using-fzf-in-your-program/
type Fzf struct {
	resultCh   chan string
	resultChan chan string

	options *fzf.Options

	// Command-line arguments for fzf, passed in the same format as the CLI.
	//
	// Example:
	// []string{"--multi", "--reverse"},
	args []string
}

// NewFzf creates a new Fzf selector instance.
//
// The provided arguments should be specified in the same way as in the CLI.
//
// Example:
// []string{"--multi", "--reverse"},
func NewFzf(args []string) (Selector, error) {
	f := &Fzf{
		args:       args,
		resultCh:   make(chan string),
		resultChan: make(chan string),
	}

	options, err := fzf.ParseOptions(true, nil)

	if err != nil {
		return nil, err
	}

	options.Output = f.resultCh
	f.options = options

	return f, nil
}

func (f *Fzf) Run(inputChan chan string) (string, error) {
	go func() {
		for out := range f.resultCh {
			f.resultChan <- out
		}

		close(f.resultChan)
	}()

	f.options.Input = inputChan

	_, err := fzf.Run(f.options)

	close(f.resultCh)

	if err != nil {
		return "", err
	}

	return <-f.resultChan, nil
}
