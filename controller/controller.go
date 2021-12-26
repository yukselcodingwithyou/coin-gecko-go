package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	httpClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
)

type Controller struct {
	controller *echo.Echo
}

func New() *Controller {
	return &Controller{
		controller: echo.New(),
	}
}

func (controller *Controller) Start(client *httpClient.Client) {
	controller.configure()
	controller.handle(client)
	controller.start(1234)
}

func (controller *Controller) configure() {
	controller.controller.HideBanner = true
	controller.controller.Logger.SetPrefix("coin-service")
	controller.controller.Logger.SetLevel(log.Lvl(4))
}

func (controller *Controller) start(port int) {
	controller.controller.Logger.Fatal(controller.controller.Start(fmt.Sprintf(":%d", port)))
}

func (controller *Controller) handle(client *httpClient.Client) {
	route(controller, client)
}
