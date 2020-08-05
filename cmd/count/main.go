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
		},

		//Commands: []cli.Command{
		//	&cli.Command{
		//		Name:  "up",
		//		Aliases: []string{"u"},
		//		Usage: "",
		//	},
		//},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
