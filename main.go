package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Health Checker",
		Usage: "A tool that checks if a website is up or down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "destination",
				Aliases: []string{"d"},
				Usage:   "Domain name to check",
			},
		},
		Action: func(c *cli.Context) error {
			status := Check(c.String("destination"))
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
