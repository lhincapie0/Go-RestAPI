package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/lhincapie0/Go-RestAPI/API/infoHandler"
	"github.com/valyala/fasthttp"
)

//GET DOMAIN INFO
func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	infoHandler.GetDomainInfo(ctx)

}

//GetSearchHistory
func GetSearchHistory(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Loading Search History")

}

//Endpoints calls
func main() {
	fmt.Println("Hello world")
	router := fasthttprouter.New()
	router.GET("/serverInfo/:server", GetDomainInfo)
	router.GET("/searchHistory/", GetSearchHistory)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
