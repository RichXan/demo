package main

import (
	"log"
	"os"
	"xproject/global"

	"github.com/urfave/cli/v2"
)

var VERSION = "0.1.0"

func main() {
	app := &cli.App{
		Name:    "seanet",
		Version: VERSION,
		Usage:   "netkit auth service",
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "/etc/seanet/seanet.yml",
			Usage:   "config file path",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "start http server",
			Action: func(c *cli.Context) error {
				global.InitConfig(c.String("config"))

				// waitgroup
				// waitgroup := &sync.WaitGroup{}
				// waitgroup.Add(1)
				// go func() {
				// 	defer waitgroup.Done()
				// 	http.Start()
				// }()
				// waitgroup.Wait()
				return nil
			},
		},
		// {
		// 	Name:  "http",
		// 	Usage: "start api gateway.",
		// 	Action: func(c *cli.Context) error {
		// 		global.InitConfig(c.String("config"))

		// 		http.Start()

		// 		return nil
		// 	},
		// },
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
