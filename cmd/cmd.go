package cmd

import (
	"github.com/iron-kit/ironic/new"
	"github.com/iron-kit/ironic/proto"
	"log"
	"os"
	// "github.com/micro/go-micro/cmd"
	"github.com/urfave/cli"
)

var (
	name        = "miron cli"
	description = "custom cloud-native toolkit based go-micro"
	version     = "0.1.0"
)

func Init() {
	// Setup(cmd.App())
	app := cli.NewApp()
	Setup(app)
	app.Name = name
	app.Description = description
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Setup(app *cli.App) {
	app.Commands = append(app.Commands, new.Commands()...)
	app.Commands = append(app.Commands, proto.Commands()...)
	app.Action = func(context *cli.Context) {
		cli.ShowAppHelp(context)
	}

	// setup(app)
}
