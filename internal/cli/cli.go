package cli

import (
	"os"
	"runtime/debug"

	"github.com/gabefiori/gosp/internal/app"
	"github.com/gabefiori/gosp/internal/config"
	"github.com/urfave/cli/v2"
)

// Run initializes and executes the command-line interface (CLI) application.
func Run() error {
	var (
		path         string
		expandOutput bool
		measure      bool
		selector     string
	)

	app := &cli.App{
		Name:        "Select Projects",
		HelpName:    "gosp",
		Usage:       "Select projects",
		Description: "A simple tool for quickly selecting projects.",
		Version:     getVersion(),

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `file`",
				Value:       "~/.config/gosp/config.json",
				TakesFile:   true,
				Destination: &path,
			},
			&cli.StringFlag{
				Name:        "selector",
				Aliases:     []string{"s"},
				Usage:       "Selector for displaying the projects (available: \"fzf\", \"fzy\")",
				Value:       "fzf",
				Destination: &selector,
			},
			&cli.BoolFlag{
				Name:        "expand-output",
				Aliases:     []string{"eo"},
				Usage:       "Expand output",
				Value:       true,
				Destination: &expandOutput,
			},
			&cli.BoolFlag{
				Name:        "measure",
				Aliases:     []string{"m"},
				Usage:       "Measure performance (time taken and number of items)",
				Value:       false,
				Destination: &measure,
			},
		},

		Action: func(ctx *cli.Context) error {
			params := &config.LoadParams{
				Path:    path,
				Measure: measure,
			}

			if ctx.IsSet("expand-output") {
				params.ExpandOutput = &expandOutput
			}

			if ctx.IsSet("selector") {
				params.Selector = selector
			}

			cfg, err := config.Load(params)

			if err != nil {
				return err
			}

			return app.Run(cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}

	return "unknown"
}
