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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com"
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search by IPs address on internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search by servers names on internet",
			Flags:  flags,
			Action: searchServers,
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

// Example of command -> go run main.go servers --host google.com
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
