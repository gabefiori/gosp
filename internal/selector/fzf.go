package selector

import (
	fzf "github.com/junegunn/fzf/src"
)

// Fzf selector for command-line fuzzy finding.
//
// For more information, see:
// https://junegunn.github.io/fzf/tips/using-fzf-in-your-program/
type Fzf struct {
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
func NewFzf(args []string) Selector {
	return &Fzf{
		args: args,
	}
}

func (f *Fzf) Run(inputChan chan string) (string, error) {
	outputChan := make(chan string)
	resultChan := make(chan string)

	go func() {
		for out := range outputChan {
			resultChan <- out
		}

		close(resultChan)
	}()

	options, err := fzf.ParseOptions(true, nil)

	if err != nil {
		return "", err
	}

	options.Input = inputChan
	options.Output = outputChan

	_, err = fzf.Run(options)

	close(outputChan)

	if err != nil {
		return "", err
	}

	return <-resultChan, nil
}
