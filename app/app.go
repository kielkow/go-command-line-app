package app

import "github/urfave/cli"

// Generate CLI app
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CommandLineApp"
	app.Usage = "Search by server names and IPs"

	return app
}
