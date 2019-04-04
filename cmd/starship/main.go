package main

import (
	"fmt"
	"log"
	"os"

	framework "github.com/multiverse-os/starshipyard/framework"
	cli "github.com/urfave/cli"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "starship"
	cmd.Usage = "A command-line tool for controling the starshipyard server, scaffolding boilerplate code, and executing developer defined commands"
	cmd.Version = "0.1.0"
	cmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "env, e",
			Value:  "development",
			Usage:  "Specify the server environment",
			EnvVar: "STARSHIP_ENV",
		},
		cli.StringFlag{
			Name:   "address, a",
			Value:  "localhost",
			Usage:  "Specify the address for the HTTP server to listen",
			EnvVar: "STARSHIP_ADDRESS",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "8080",
			Usage:  "Specify the listening port for the HTTP server",
			EnvVar: "STARSHIP_ADDRESS",
		},
		//cli.BoolFlag{
		//	Name:   "daemon, d",
		//	Usage:  "Daemonize the http server",
		//	EnvVar: "STARSHIP_ADDRESS",
		//},
	}
	cmd.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Options for controlling starshipyard HTTP server",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "Start the starship yard http server",
					Action: func(c *cli.Context) error {
						// TODO: Use flags to get port and host address and environment to
						// start the server in or use envirnonemtnal variables. We take
						// these and overwrite the values in the config object in this
						// function
						config, err := framework.LoadConfig("config/app.yaml")
						if err != nil {
							fmt.Println("[starship] missing 'config/app.yaml' starship app configuration")
							config = framework.DefaultConfig()
						}
						fmt.Println("Calling framework.Init(config)")
						s := framework.Init(config)
						fmt.Println("framework.Init(config) completed")

						//if c.Bool("daemon") {
						//	fmt.Println("[starship] launching in daemon mode...")
						//	s.StartAsDaemon()
						//} else {

						//fmt.Println("app.Process.Pid:", s.Process.Pid)
						//fmt.Println("app.Process.Path:", s.Process.Path)
						fmt.Println("calling framework.Start()")
						s.Start()
						fmt.Println("framework.Start() completed")
						//}
						return nil
					},
				},
			},
		},
		//{
		//	Name:    "generate",
		//	Aliases: []string{"g"},
		//	Usage:   "Generate new go source code for models, controllers, and views",
		//	Subcommands: []cli.Command{
		//		{
		//			Name:  "model",
		//			Usage: "Build a model template with the specified object data",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("[starship] code generation functionality is not implemented yet")
		//				fmt.Println("[starship] test code has been built, it just needs to be migrated into the base and will be available shortly")
		//				return nil
		//			},
		//		},
		//		{
		//			Name:  "controller",
		//			Usage: "Build a controller template with the specified object data",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("[starship] code generation functionality is not implemented yet")
		//				fmt.Println("[starship] test code has been built, it just needs to be migrated into the base and will be available shortly")
		//				return nil
		//			},
		//		},
		//		{
		//			Name:  "view",
		//			Usage: "Build a view template with the specified object data",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("[starship] code generation functionality is not implemented yet")
		//				fmt.Println("[starship] test code has been built, it just needs to be migrated into the base and will be available shortly")
		//				return nil
		//			},
		//		},
		//		{
		//			Name:  "job",
		//			Usage: "Build a job template with the specified object data",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("[starship] code generation functionality is not implemented yet")
		//				fmt.Println("[starship] test code has been built, it just needs to be migrated into the base and will be available shortly")
		//				return nil
		//			},
		//		},
		//		{
		//			Name:  "model",
		//			Usage: "Build a model template with the specified object data",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("[starship] code generation functionality is not implemented yet")
		//				fmt.Println("[starship] test code has been built, it just needs to be migrated into the base and will be available shortly")
		//				return nil
		//			},
		//		},
		//	},
		//},
		//{
		//	Name:    "console",
		//	Aliases: []string{"c"},
		//	Usage:   "Start the starship yard console interface",
		//	Action: func(c *cli.Context) error {
		//		fmt.Println("[starship][CONSOLE] console interface is not implemented yes")
		//		return nil
		//	},
		//},
	}
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
