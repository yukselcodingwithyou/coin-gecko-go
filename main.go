package main

import "github.com/yukselcodingwithyou/gocoingecko/application"

func main() {
	application.New().Start("mongodb+srv://username:password123!.@cluster0.rlkfg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
}
