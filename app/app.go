package app

import (
	"log"
	"net"
	"fmt"

	"github/urfave/cli"
)

// Generate CLI app
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CommandLineApp"
	app.Usage = "Search by server names and IPs"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search by IPs address on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com"
				}
			},
			Action: searchIps,
		},
	}

	return app
}

// Example of command -> go run main.go ip --host google.com
func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
