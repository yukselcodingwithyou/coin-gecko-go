package cli

import (
	"github.com/urfave/cli/v2"
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	hClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
	mClient "github.com/yukselcodingwithyou/gocoingecko/client/mongo"
	"github.com/yukselcodingwithyou/gocoingecko/controller"
	"github.com/yukselcodingwithyou/gocoingecko/handler"
)

type Client struct {
	Application *cli.App
}

func New(httpClient *hClient.Client, mongoClient *mClient.Client) *Client {
	return &Client{
		Application: &cli.App{
			Name:  "Go Gecko Go",
			Usage: "An app implements gecko coin methods",
			Commands: []*cli.Command{
				{
					Name:    "start-http-server",
					Aliases: []string{"s"},
					Usage:   "starts echo server",
					Action: func(c *cli.Context) error {
						controller.New().Start(httpClient)
						return nil
					},
				},
				{
					Name:    "ping",
					Aliases: []string{"p"},
					Usage:   "ping gecko api",
					Action: func(c *cli.Context) error {
						handler.Get(fromClient.Ping).Handle(httpClient, mongoClient, fromClient.CLI)
						return nil
					},
				},
				{
					Name:    "list-coins",
					Aliases: []string{"lc"},
					Usage:   "list all available coins",
					Action: func(c *cli.Context) error {
						handler.Get(fromClient.ListCoins).Handle(httpClient, mongoClient, fromClient.CLI)
						return nil
					},
				},
			},
		},
	}
}
