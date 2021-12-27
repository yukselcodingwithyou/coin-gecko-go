package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	hClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
	mClient "github.com/yukselcodingwithyou/gocoingecko/client/mongo"
	"log"
	"net/http"
)

type pingHandler struct {
}

func newPingHandler() *pingHandler {
	return new(pingHandler)
}

func (ph *pingHandler) Handle(httpClient *hClient.Client, mongoClient *mClient.Client, clientType fromClient.FromClient) echo.HandlerFunc {
	ping, err := httpClient.Ping()
	if err != nil {
		log.Println(err)
		return getResponse(http.StatusGone, ping)
	}
	if fromClient.NewFromClient().IsFromCli(clientType) {
		fmt.Println(ping.GeckoSays)
	}
	return getResponse(http.StatusOK, *ping)
}
