package main

import (
	"context"
	"errors"
	"git.cyradar.com/phinc/my-awesome-project/internal/app/server"
	"git.cyradar.com/phinc/my-awesome-project/pkg/util"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)


func main() {
	app := cli.NewApp()
	app.Name = "Staff Management"
	app.Usage = "Demo Microservice"
	app.Authors = []*cli.Author{
		{
			Name:  "Quang Nguyen",
			Email: "quangnc@cyradar.com",
		},
		{
			Name:  "Chi Nguyen",
			Email: "chint@cyradar.com",
		},
	}

	app.Copyright = "Copyright © 2020 CyRadar. All Rights Reserved."
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Commands = []*cli.Command{
		{
			Name:   "serve",
			Usage:  "serve as an api server",
			Action: Serve(),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "addr",
					Value: "localhost:8080",
					Usage: "specify address to serve on",
				},

				&cli.StringFlag{
					Name:  "env",
					Value: "./.env",
					Usage: "specify .env file",
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		err := util.LoadEnv(c.String("env"))
		if err != nil {
			log.Fatalln(err)
		}

		app.Metadata["context"] = context.Background()
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err == nil {
		return
	}

	log.Println(err)
	os.Exit(1)
}

func Serve() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		rootCtx, ok := c.App.Metadata["context"].(context.Context)
		if !ok {
			return errors.New("invalid root context")
		}

		addr := c.String("addr")

		return server.ServeServer(rootCtx, addr)
	}
}
