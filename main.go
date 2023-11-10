package main

import (
	"log"
	"os"

	"github.com/furtidev/tplinkctl/wr840n"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "tplinkctl",
		Usage: "control your TP-Link router from the command line.",
		UsageText: "tplinkctl [ROUTER MODEL] --user <username> --pass <password>",
		Version: "0.1",
		Commands: []*cli.Command{
			{
				Name: "wr840n",
				Usage: "control your TL-WR840N router.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "user",
						Value: "admin",
						Usage: "username to login to your panel",
					},
					&cli.StringFlag{
						Name: "pass",
						Value: "admin",
						Usage: "password to login to your panel",
					},
				},
				Before: wr840n.Setup,
				Action: wr840n.Status,
				Subcommands: []*cli.Command {
					{
						Name: "clients",
						Usage: "check connected DHCP clients.",
						Action: wr840n.Clients,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err)
	}
}