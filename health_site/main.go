package main

import (
	"fmt"
	"os"
	//"net/http"
	"log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App {
		Name: "HealthCheck",
		Usage: "Checking website down or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
                Name:     "domain",
                Aliases:  []string{"d"},
                Required: true,
                Usage:   "URL to check",
            },
			&cli.StringFlag{
				Name: "port",
				Aliases:  []string{"p"},
                //Value:    "80",
                Usage:   "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String ("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}