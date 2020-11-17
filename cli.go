package main

import (
	"github.com/ggdream/ngo/app"
	"github.com/ggdream/ngo/conf"
	"github.com/urfave/cli/v2"
	"os"
)


func runCli() error {
	cApp := cli.NewApp()
	cApp.Name = conf.NAME
	cApp.Version = conf.VERSION
	cApp.Usage = conf.USAGE


	cApp.HideHelp = true
	cApp.Flags = []cli.Flag{
		fHost, fPort, fBind, fStatic, fDaemon, fLog,
	}
	cApp.Action = func(c *cli.Context) error {
		return app.New()
	}

	return cApp.Run(os.Args)
}



var fHost = &cli.StringFlag{
	Name: "host",
	Aliases: []string{"h"},
	Usage: "Specify host.",
	Value: conf.HOST,
	Destination: &conf.HOST,
}

var fPort = &cli.IntFlag{
	Name: "port",
	Aliases: []string{"p"},
	Usage: "Specify port.",
	Value: conf.PORT,
	Destination: &conf.PORT,
}

var fBind = &cli.StringFlag{
	Name: "bind",
	Aliases: []string{"b"},
	Usage: "Specify the route to which you want to bind.",
	Value: conf.PATH,
	Destination: &conf.PATH,
}

var fStatic = &cli.StringFlag{
	Name: "static",
	Aliases: []string{"s"},
	Usage: "Specify the static files path.",
	Value: conf.STATIC,
	Destination: &conf.STATIC,
}

var fDaemon = &cli.BoolFlag{
	Name: "daemon",
	Aliases: []string{"d"},
	Usage: "Start the daemon.",
	Value: conf.Daemon,
	Destination: &conf.Daemon,
}

var fLog = &cli.BoolFlag{
	Name: "log",
	Aliases: []string{"l"},
	Usage: "Close the logger.",
	Value: conf.PrintLog,
	Destination: &conf.PrintLog,
}
