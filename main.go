package main

import (
	"encoding/json"
	"log"
	"os"
	"path"

	cli "github.com/helderfarias/dynamock/cli"

	settings "github.com/codegangsta/cli"
)

func main() {
	app := settings.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "Dynamic Mock for Api Rest"
	app.Version = "1.0.0"

	app.Flags = []settings.Flag{
		settings.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "-c config.json",
		},
		settings.StringFlag{
			Name:  "port, p",
			Value: "3010",
			Usage: "-p 3010",
		},
	}

	app.Action = func(ctx *settings.Context) error {
		configFile, err := os.Open(ctx.GlobalString("config"))
		defer configFile.Close()

		if err != nil {
			log.Println("opening config file", err.Error())
		}

		config := cli.Configuration{}
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&config); err != nil {
			log.Println("parsing config file", err.Error())
		}

		cli.Run(&config)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
