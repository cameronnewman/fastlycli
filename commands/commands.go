package commands

import (
	"github.com/cameronnewman/fastlycli/fastlyclient"
	"github.com/codegangsta/cli"
	"os"
)

func Excute() {

	app := cli.NewApp()
	app.Name = "fastlycli"
	app.Usage = "Manage Fastly CDN Services via the cli"
	app.Version = "0.9.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "enable verbose logging",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "service",
			Usage: "Get Service Details",
			Subcommands: []cli.Command{
				{
					Name:  "details",
					Usage: "Gets service details",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "service, s",
							Usage: "Service Name defined in app.fastly.com",
						},
					},
					Action: func(c *cli.Context) {
						if c.String("service") != "" {
							fastly := fastlyclient.NewFastlyClient()
							fastly.GetServiceDetails(c.String("service"))
						}
					},
				},
				{
					Name:  "domains",
					Usage: "Gets a service domains",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "service, s",
							Usage: "Service Name defined in app.fastly.com",
						},
					},
					Action: func(c *cli.Context) {
						if c.String("service") != "" {
							fastly := fastlyclient.NewFastlyClient()
							fastly.GetServiceDomains(c.String("service"))
						}
					},
				},

				{
					Name:  "backends",
					Usage: "Gets a service domains",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "service, s",
							Usage: "Service Name defined in app.fastly.com",
						},
					},
					Action: func(c *cli.Context) {
						if c.String("service") != "" {
							fastly := fastlyclient.NewFastlyClient()
							fastly.GetServiceBackends(c.String("service"))
						}
					},
				},
			},
		},
		{
			Name:  "purge",
			Usage: "Purge objects from the CDN",
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
			Action: func(c *cli.Context) {
				if c.IsSet("service") && c.String("service") != "" {
					fastly := fastlyclient.NewFastlyClient()
					fastly.PurgeObjects(c.String("service"), c.String("object"))
				} else {
					println("No service name defined")
				}
			},
		},
	}

	app.Run(os.Args)
}
