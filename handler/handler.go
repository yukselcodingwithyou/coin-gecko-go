package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	fromClient "github.com/yukselcodingwithyou/gocoingecko/client"
	hClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
	mClient "github.com/yukselcodingwithyou/gocoingecko/client/mongo"
)

type Handler interface {
	Handle(httpClient *hClient.Client, mongoClient *mClient.Client, clientType fromClient.FromClient) echo.HandlerFunc
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

func getResponse(status int, body interface{}) func(context echo.Context) error {
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		context.Response().WriteHeader(status)
		return json.NewEncoder(context.Response()).Encode(body)
	}
}
