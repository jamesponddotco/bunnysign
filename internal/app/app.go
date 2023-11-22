// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"
	"time"

	"git.sr.ht/~jamesponddotco/bunnysign/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Usage:    "the URL to sign",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "key",
			Aliases:  []string{"k"},
			Usage:    "the key to sign with",
			Required: true,
		},
		&cli.DurationFlag{
			Name:    "expiration",
			Aliases: []string{"e"},
			Usage:   "the expiration time of the signature",
			Value:   365 * 24 * time.Hour,
		},
	}

	app.Action = SignAction

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
