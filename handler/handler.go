package handler

import (
	"github.com/labstack/echo/v4"
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	httpClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
)

type Handler interface {
	Handle(client *httpClient.Client, clientType fromClient.FromClient) echo.HandlerFunc
}

func Get(methodName fromClient.ImplementedMethodName) Handler {
	switch methodName {
	case fromClient.Ping:
		return newPingHandler()
	case fromClient.ListCoins:
		return newListCoinsHandler()
	default:
		return nil
	}
}
