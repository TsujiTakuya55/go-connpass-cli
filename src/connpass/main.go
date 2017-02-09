package main

import (
	"connpass/cmd"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var connpass *cmd.Connpass

func main() {
	app := cli.NewApp()
	app.Version = "0.1"
	app.Name = "go-conpass-cli"
	app.Usage = "conpass --help"

	globalFlags := []cli.Flag{
		cli.StringFlag{
			Name:"count,c",
			Value:"10",
			Usage:"",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:"search",
			Aliases:[]string{"s"},
			Usage:"Search event information with any word",
			Flags:append([]cli.Flag{
				cli.StringFlag{
					Name:"keyWord,k",
					Value:"",
					Usage:"",
				},
			}, globalFlags...),
			//Before: ,
			Action:func(c *cli.Context) {
				cmd.Search(connpass, c.String("keyWord"), c.String("count"))
			},
		},
		{
			Name:"list",
			Aliases:[]string{"l"},
			Usage:"Displayed in order of opening date and time (default 10)",
			Flags:append([]cli.Flag{},globalFlags...),
			Action:func(c *cli.Context) {
				cmd.List(connpass, c.String("count"))
			},
		},
	}

	app.Run(os.Args)
}