package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	httpClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
	"log"
	"net/http"
)

type listCoinsHandler struct {
}

func newListCoinsHandler() *listCoinsHandler {
	return new(listCoinsHandler)
}

func (lch *listCoinsHandler) Handle(client *httpClient.Client, clientType fromClient.FromClient) echo.HandlerFunc {
	list, err := client.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	if fromClient.NewFromClient().IsFromCli(clientType) {
		fmt.Println("Available coins in total: ", len(*list))
		fmt.Println(*list)
	}
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		context.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(context.Response()).Encode(*list)
	}
}
