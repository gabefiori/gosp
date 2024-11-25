package config

import (
	"os"

	"github.com/gabefiori/gosp/internal/finder"
	"github.com/goccy/go-json"
	"github.com/mitchellh/go-homedir"
)

// Config represents the configuration structure for the application.
type Config struct {
	// List of sources to be used by the finder
	Sources []finder.Source `json:"sources"`

	// Flag to indicate if output should be expanded
	// Useful to hide the user's home directory
	ExpandOutput bool `json:"expand_output"`

	// Flag to indicate if measurement should be performed
	Measure bool

	// Selector for displaying the projects
	Selector string `json:"selector"`
}

type LoadParams struct {
	Path         string
	Measure      bool
	ExpandOutput *bool
	Selector     string
}

// Load reads the configuration from a JSON file at the specified path.
func Load(params *LoadParams) (*Config, error) {
	path, err := homedir.Expand(params.Path)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg Config

	cfg.ExpandOutput = true

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	cfg.Measure = params.Measure

	if params.ExpandOutput != nil {
		cfg.ExpandOutput = *params.ExpandOutput
	}

	if params.Selector != "" {
		cfg.Selector = params.Selector
	}

	if cfg.Selector == "" {
		cfg.Selector = "fzf"
	}

	return &cfg, nil
}
