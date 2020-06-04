package main

import (
	"database/sql"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/lhincapie0/Go-RestAPI/API/database"
	"github.com/lhincapie0/Go-RestAPI/API/infohandler"
	"github.com/valyala/fasthttp"
)

//GET DOMAIN INFO
func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	infohandler.GetDomainInfo(ctx)

}

//GetSearchHistory
func GetSearchHistory(ctx *fasthttp.RequestCtx) {
	infohandler.GetDomainsHistory(ctx)
}

func startDB() {
	var db *sql.DB
	db = database.ConnectDB()
	infohandler.HTTPInfoHandler(db)
}

func main() {
	router := fasthttprouter.New()
	startDB()

	//Endpoints calls
	router.GET("/serverInfo/:server", GetDomainInfo)
	router.GET("/searchHistory/", GetSearchHistory)

	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
}
