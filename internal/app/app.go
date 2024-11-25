package app

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gabefiori/gosp/internal/config"
	"github.com/gabefiori/gosp/internal/finder"
	"github.com/gabefiori/gosp/internal/selector"
	"github.com/mitchellh/go-homedir"
)

// Run executes the main logic of the application using the provided configuration.
func Run(cfg *config.Config) error {
	home, err := homedir.Dir()

	if err != nil {
		return err
	}

	// Channel to receive output (string) from the finder.
	//
	// This channel is also passed to the selector to populate its input.
	outputChan := make(chan string)

	measureStart := time.Now()
	buf := bytes.NewBufferString("")

	go finder.Run(&finder.FinderOpts{
		Sources:    cfg.Sources,
		OutputChan: outputChan,
		HomeDir:    home,
	})

	// If output expansion is not enabled, set the home directory to "~".
	// This is useful for hiding the user's home directory.
	if !cfg.ExpandOutput {
		home = "~"
	}

	// If measurement is enabled, count the number of projects found
	// and the time taken to find the projects.
	if cfg.Measure {
		var count int

		for range outputChan {
			count++
		}

		measureEnd := time.Since(measureStart).String()
		msg := fmt.Sprintf("Took %s (%d projects)", measureEnd, count)

		if _, err := buf.WriteString(msg); err != nil {
			return err
		}

		_, err = io.Copy(os.Stdout, buf)
		return err
	}

	t := selector.TypeFromStr(cfg.Selector)
	s := selector.New(t, nil)

	result, err := s.Run(outputChan)

	// An empty result indicates that the selector was canceled.
	if err != nil || result == "" {
		return err
	}

	// The first character ("~") of the result is skipped.
	// It's only used for display inside the selector.
	//
	// The expanded version of the result must be used;
	// otherwise, it will not be able to be consumed by other programs.
	if _, err := buf.WriteString(home + result[1:]); err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, buf)
	return err
}
