package application

import (
	cliClient "github.com/yukselcodingwithyou/gocoingecko/client/cli"
	httpClient "github.com/yukselcodingwithyou/gocoingecko/client/http"
	mongoClient "github.com/yukselcodingwithyou/gocoingecko/client/mongo"
	"log"
	"net/http"
	"os"
	"time"
)

var MongoClient *mongoClient.Client
var HttpClient *httpClient.Client
var CliClient *cliClient.Client

type Application struct {
}

func New() *Application {
	return new(Application)
}

func (a *Application) Start(uri string) {
	startMongoClient(uri)
	startHttpClient()
	startCliClient()
}

func startMongoClient(uri string) {
	MongoClient = mongoClient.New(uri).Connect()
}

func startHttpClient() {
	HttpClient = httpClient.New(&http.Client{
		Timeout: time.Second * 10,
	})
}

func startCliClient() {
	CliClient = cliClient.New(HttpClient)
	err := CliClient.Application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
