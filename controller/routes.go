package controller

import (
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	"github.com/yukselcodingwithyou/gocoingecko/client/http"
	"github.com/yukselcodingwithyou/gocoingecko/handler"
)

const (
	PingPath         = "/ping"
	ListAllCoinsPath = "/coins"
)

func route(controller *Controller, client *http.Client) {
	controller.controller.GET(PingPath, handler.Get(fromClient.Ping).Handle(client, fromClient.HTTP))
	controller.controller.GET(ListAllCoinsPath, handler.Get(fromClient.ListCoins).Handle(client, fromClient.HTTP))
}
