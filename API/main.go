package main

import (
	"database/sql"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/lhincapie0/Go-RestAPI/API/database"
	"github.com/lhincapie0/Go-RestAPI/API/infoHandler"
	"github.com/valyala/fasthttp"
)

//GET DOMAIN INFO
func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	infoHandler.GetDomainInfo(ctx)

}

//GetSearchHistory
func GetSearchHistory(ctx *fasthttp.RequestCtx) {
	infoHandler.GetDomainsHistory(ctx)
	//	infoHandler.getDomainsHistory()
}

func startDB() {
	var b *sql.DB
	b = database.ConnectDB()
	infoHandler.HttpInfoHandler(b)
}

//Endpoints calls
func main() {
	router := fasthttprouter.New()
	router.GET("/serverInfo/:server", GetDomainInfo)
	router.GET("/searchHistory/", GetSearchHistory)
	startDB()

	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
}
