package main

import (
	"github.com/BUGLAN/zk-batch/batch"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func buildApp() *cli.App {
	app := &cli.App{
		Name: "zku",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server",
				Aliases: []string{"s"},
				Value:   "localhost:2181",
				Usage:   "specify zookeeper server",
			},
			&cli.StringFlag{
				Name:    "auth",
				Aliases: []string{"a"},
				Usage:   "zookeeper auth passport",
			},
			&cli.StringFlag{
				Name:    "digest",
				Aliases: []string{"u"},
				Value:   "digest",
				Usage:   "zookeeper auth digest",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "import",
				Aliases: []string{"i"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "filename",
						Aliases:  []string{"f"},
						Usage:    "指定导入的文件",
						Required: true,
					},
				},
				Usage:  "import data to zookeeper",
				Action: batch.Import,
			},
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "export data to zookeeper",
				Action:  batch.Export,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "filename",
						Aliases: []string{"f"},
						Usage:   "到处到指定的文件",
					},
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "root path",
						Value:   "/",
					},
				},
			},
		},
	}

	// sort command and flags
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	return app
}

func main() {
	app := buildApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
