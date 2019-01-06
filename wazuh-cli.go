package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrtc0/wazuh"
	"github.com/urfave/cli"

	. "github.com/mrtc0/wazuh-cli/command"
)

var client *wazuh.Client

func init() {
	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client = wazuh.New(config.APIURL)
	client.SetBasicAuth(config.BasicAuth.Username, config.BasicAuth.Password)
}

func main() {
	app := cli.NewApp()
	app.Name = "wazuh-cli"
	app.Usage = "manage Wazuh REST API resources"

	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "show config",
			Subcommands: []cli.Command{
				{
					Name:  "show",
					Usage: "Get config",
					Action: func(c *cli.Context) error {
						config, err := GetConfig()
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						ShowConfig(*config)
						return nil
					},
				},
			},
		},
		{
			Name:    "agent",
			Aliases: []string{"a"},
			Usage:   "options for agent operations",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show all agents",
					Action: func(c *cli.Context) error {
						GetAllAgents(client)
						return nil
					},
				},
				{
					Name:  "get",
					Usage: "show information an agents",
					Action: func(c *cli.Context) error {
						fmt.Println("wip ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:    "syscollector",
			Aliases: []string{"s"},
			Usage:   "options for syscollector operations",
			Subcommands: []cli.Command{
				{
					Name:  "hardware",
					Usage: "show hardware information",
					Action: func(c *cli.Context) error {
						GetHardwareInfo(client, c.Args().First())
						return nil
					},
				},
				{
					Name:  "address",
					Usage: "show network address information",
					Action: func(c *cli.Context) error {
						GetNetworkAddrInfo(client, c.Args().First())
						return nil
					},
				},
				{
					Name:  "packages",
					Usage: "show all packages",
					Action: func(c *cli.Context) error {
						GetPackagesInfo(client, c.Args().First())
						return nil
					},
				},
				{
					Name:  "ports",
					Usage: "show all listen ports",
					Action: func(c *cli.Context) error {
						GetOpenPorts(client, c.Args().First())
						return nil
					},
				},
				{
					Name:  "processes",
					Usage: "show all processes",
					Action: func(c *cli.Context) error {
						GetProcesses(client, c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
