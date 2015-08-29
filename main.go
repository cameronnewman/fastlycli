package main

import (
	"github.com/cameronnewman/fastlycli/fastlyservice"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "fastlycli"
	app.EnableBashCompletion = true
	app.Usage = "Manage Fastly CDN Services via the cli"
	app.Version = "0.5.0"
	app.Commands = []cli.Command{
		{
			Name:    "purge",
			Aliases: []string{"p"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "service, s",
					Usage: "Service Name defined in app.fastly.com",
				},
				cli.StringFlag{
					Name:  "object, o",
					Usage: "Objects to be purged",
				},
			},
			Usage: "Purge objects from the CDN",
			Action: func(c *cli.Context) {
				if c.String("service") != "" {
					fastly := fastlyservice.NewFastlyService()
					fastly.PurgeObjects(c.String("service"), c.String("object"))
				} else {
					println("No CDN service name defined")
				}
			},
		},
		{
			Name:    "service",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "service, s",
					Usage: "Service Name defined in app.fastly.com",
				},
			},
			Usage: "Get Service Details",
			Action: func(c *cli.Context) {
				if c.String("service") != "" {
					fastly := fastlyservice.NewFastlyService()
					fastly.ReturnServiceDetails(c.String("service"))
				} else {
					println("No CDN service name defined")
				}
			},
		},
	}

	app.Run(os.Args)
}
