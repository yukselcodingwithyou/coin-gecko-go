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

type listCoinsHandler struct {
}

func newListCoinsHandler() *listCoinsHandler {
	return new(listCoinsHandler)
}

func (lch *listCoinsHandler) Handle(httpClient *hClient.Client, mongoClient *mClient.Client, clientType fromClient.FromClient) echo.HandlerFunc {
	list, err := httpClient.CoinsList()
	if err != nil {
		log.Println(err)
		return getResponse(http.StatusGone, list)
	}
	if fromClient.NewFromClient().IsFromCli(clientType) {
		fmt.Println("Available coins in total: ", len(*list))
		fmt.Println(*list)
	}
	return getResponse(http.StatusOK, *list)
}
