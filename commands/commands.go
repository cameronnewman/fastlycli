package commands

import (
	"fmt"
	"os"

	"github.com/cameronnewman/fastlycli/fastly"
	"github.com/urfave/cli"
)

//Excute main section
func Excute() {

	app := cli.NewApp()
	app.Name = "fastly"
	app.Usage = "CLI to manage Fastly CDN Services"
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
							print(fastly.New().GetService(c.String("service")))
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
							print(fastly.New().GetServiceDomains(c.String("service")))
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
							print(fastly.New().GetServiceBackends(c.String("service")))
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
					print(fastly.New().PurgeObject(c.String("service"), c.String("object")))
				} else {
					println("No service name defined")
				}
			},
		},
		{
			Name:  "purgeall",
			Usage: "Purge all objects from the CDN",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "service, s",
					Usage: "Service Name defined in app.fastly.com",
				},
			},
			Action: func(c *cli.Context) {
				if c.IsSet("service") && c.String("service") != "" {
					print(fastly.New().PurgeAllObjects(c.String("service")))
				} else {
					println("No service name defined")
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func print(r string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
