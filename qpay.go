package main

import (
	"fmt"
	"log"
	"os"

	commands "github.com/sproule-sQuad/qpcli/lib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "qpay",
		Commands: []*cli.Command{
			{
				Name:    "btc",
				Aliases: []string{"bp"},
				Usage:   "checks bitcoin price",
				Action:  func(c *cli.Context) error {
					commands.FetchBtcData()
					return nil
				},
			},
			{
				Name:    "install-dev",
				Usage:   "setup initial dev",
				Action:  func(c *cli.Context) error {
					commands.InstallDev()
					return nil
				},
			},
			{
				Name:        "azure",
				Aliases:     []string{"az"},
				Usage:       "useful azure commands bundled",
				Subcommands: []*cli.Command{
					{
						Name:  "kubectxs",
						Usage: "fetch all kubernetes contexts",
						Aliases: []string{"kctxs"},
						Action: func(c *cli.Context) error {
							fmt.Println("fetch all kuebernetes")
							return nil
						},
					},
//					{
//						Name:  "remove",
//						Usage: "remove an existing template",
//						Action: func(c *cli.Context) error {
//							fmt.Println("removed task template: ", c.Args().First())
//							return nil
//						},
//					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}