package main

import (
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/generator"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

const (
	projectPathFlag    = "project-path"
	configPathFlag     = "config"
	noInteractModeFlag = "no-interact"

	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Name:                   "skeleton",
		Usage:                  "A simple tool for generating skeleton",
		UseShortOptionHandling: true,
		Version:                version,
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize skeleton",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    projectPathFlag,
						Usage:   "project path",
						Value:   ".",
						Aliases: []string{"p"},
					},
					&cli.StringFlag{
						Name:    configPathFlag,
						Aliases: []string{"c"},
						Usage:   "config file",
						Value:   "skeleton.json",
					},
					&cli.BoolFlag{
						Name:  noInteractModeFlag,
						Usage: "disable interactive mode",
						Value: false,
					},
				},
				Action: cmdInit,
			},
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdInit(c *cli.Context) error {
	dir := c.String(projectPathFlag)
	configPath := c.String(configPathFlag)
	interactMode := !c.Bool(noInteractModeFlag)

	return generator.NewGenerator().Generate(c.Context, dir, configPath, interactMode)
}
