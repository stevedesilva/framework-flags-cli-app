package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	counter()
}

func counter() {
	app := &cli.App{
		Name:    "Counter",
		Usage:   "Count up or down",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "up",
				Aliases: []string{"u"},
				Usage:   "Count up",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "stop",
						Aliases: []string{"s"},
						Usage:   "value to count up to",
						Value:   10,
					},
				},
				Action: func(c *cli.Context) error {
					stop := c.Int("stop")
					fmt.Println("Stop =", stop)
					if stop < 0 {
						fmt.Println("Stop cannot be negative.")
					}
					for i := 1; i <= stop; i++ {
						fmt.Println(i)
					}
					return nil
				},
			},
			{
				Name:    "down",
				Aliases: []string{"d"},
				Usage:   "Count down",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "start",
						Aliases: []string{"s"},
						Usage:   "Start counting down from",
						Value:   10,
					},
				},
				Action: func(c *cli.Context) error {
					start := c.Int("start")
					if start < 0 {
						fmt.Println("Start cannot be negative")
					}
					for i := start; i >= 0; i-- {
						fmt.Println(i)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
