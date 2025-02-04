package main

import (
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/generator"
	"github.com/Dmitrijlin/go-skeleton/internal/initializer"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

const (
	projectPathFlag = "project-path"
	configPathFlag  = "config"
	//noInteractModeFlag = "no-interact"
	globalFlag = "global"

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
					&cli.BoolFlag{
						Name:    globalFlag,
						Aliases: []string{"g"},
						Usage:   "initialize skeleton in the home directory",
					},
					&cli.StringFlag{
						Name:    projectPathFlag,
						Usage:   "project path",
						Aliases: []string{"p"},
					},
					&cli.StringFlag{
						Name:    configPathFlag,
						Aliases: []string{"c"},
						Usage:   "config file",
					},
				},
				Action: cmdInit,
			},
			{
				Name:  "generate",
				Usage: "generate skeleton",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    projectPathFlag,
						Usage:   "project path",
						Aliases: []string{"p"},
					},
				},
				Action: cmdGenerate,
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
	global := c.Bool(globalFlag)
	projectPath := c.String(projectPathFlag)
	configPath := c.String(configPathFlag)

	return initializer.NewInitializer(global, projectPath, configPath).Initialize(c.Context)
}

func cmdGenerate(c *cli.Context) error {
	dir := c.String(projectPathFlag)

	return generator.NewGenerator().Generate(c.Context, dir)
}
