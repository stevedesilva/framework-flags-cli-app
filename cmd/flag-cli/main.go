package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	greeter()
}

func greeter() {
	// collect flags either by passing a var (e.g. city, or by calling c.String in the Action func)
	var city string
	app := &cli.App{
		Name:        "hello-cli",
		Usage:       "Print Hello world for our developer friends",
		Description: "A greeting application for developers",
		Version:     "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "Guest",
				Usage:   "Who to say hello to",
			},
			&cli.StringFlag{
				Name:        "city",
				Aliases:     []string{"c"},
				Usage:       "Location of user",
				EnvVars:     nil,
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       "London",
				DefaultText: "Defaults to the capital of the U.K.",
				Destination: &city,
				HasBeenSet:  false,
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			fmt.Printf("Hi %s from %s!\n", name, city)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
