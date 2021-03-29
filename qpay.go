package main

import (
	"fmt"
	"log"
	"os"

	commands "github.com/sproule-sQuad/cli/lib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "qpay",
		Commands: []*cli.Command{
			{
				Name:    "btc",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action:  func(c *cli.Context) error {
					commands.Fetch_btc_data()
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action:  func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
//			{
//				Name:        "template",
//				Aliases:     []string{"t"},
//				Usage:       "options for task templates",
//				Subcommands: []*cli.Command{
//					{
//						Name:  "add",
//						Usage: "add a new template",
//						Action: func(c *cli.Context) error {
//							fmt.Println("new task template: ", c.Args().First())
//							return nil
//						},
//					},
//					{
//						Name:  "remove",
//						Usage: "remove an existing template",
//						Action: func(c *cli.Context) error {
//							fmt.Println("removed task template: ", c.Args().First())
//							return nil
//						},
//					},
//				},
//			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}