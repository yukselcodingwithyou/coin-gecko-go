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

type pingHandler struct {
}

func newPingHandler() *pingHandler {
	return new(pingHandler)
}

func (ph *pingHandler) Handle(client *httpClient.Client, clientType fromClient.FromClient) echo.HandlerFunc {
	ping, err := client.Ping()
	if err != nil {
		log.Println(err)
		return func(context echo.Context) error {
			context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			context.Response().WriteHeader(http.StatusGone)
			return json.NewEncoder(context.Response()).Encode(*ping)
		}
	}

	if fromClient.NewFromClient().IsFromCli(clientType) {
		fmt.Println(ping.GeckoSays)
	}
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		context.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(context.Response()).Encode(*ping)
	}
}
