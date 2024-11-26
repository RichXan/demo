package main

import (
	"log"
	"os"
	"x-project/global"
	"x-project/internal/http"

	"github.com/urfave/cli/v2"
)

var VERSION = "0.1.0"

func main() {
	app := &cli.App{
		Name:    "x-project",
		Version: VERSION,
		Usage:   "x-project service",
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "/etc/x-project/config.yml",
			Usage:   "config file path",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "start http server",
			Action: func(c *cli.Context) error {
				global.InitConfig(c.String("config"))
				http.Start()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
