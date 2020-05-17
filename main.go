package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/likexian/whois-go"
	"github.com/valyala/fasthttp"
)

//CONSTANTS
const serverInfoPath string = "https://api.ssllabs.com/api/v3/analyze?host="
const COUNTRY string = "Country"
const ORGANIZATION string = "Organization"

var domain Domain

type Domain struct {
	EndPoints []EndPoint `json:"endpoints"`
	Host      string     `json:"host"`
}

type DomainInfo struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}

type EndPoint struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

//PACKAGE GET DATA
func ConsumeSSLApi(server interface{}) {
	fmt.Println("Getting api")
	serverName := getStringInterface(server)
	path := serverInfoPath + serverName
	response, err := http.Get(path)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &domain)

		HandleServerInfo()

	}
}

func HandleServerInfo() []Server {

	//CREATE SERVER ARRAY
	var servers []Server
	for i := range domain.EndPoints {
		country, owner := WhoIs(domain.EndPoints[i].IpAddress)
		s := Server{
			IpAddress: domain.EndPoints[i].IpAddress,
			Grade:     domain.EndPoints[i].Grade,
			Country:   country,
			Owner:     owner,
		}
		servers = append(servers, s)
	}
	return servers
}

func WhoIs(ip string) (string, string) {
	result, err := whois.Whois(ip)
	if err == nil {
		data := strings.Split(result, "\n")

		indexCountry := findIndex(data, COUNTRY)
		country := findValueWhoIs(data, indexCountry)
		fmt.Println(country)
		indexOrganization := findIndex(data, ORGANIZATION)
		organization := findValueWhoIs(data, indexOrganization)

		fmt.Println(organization)
		return country, organization

	}
	return "", ""
}

func findValueWhoIs(data []string, index int) string {
	val := strings.Split(data[index], ":")
	val2 := strings.Split(val[1], "  ")
	for i := range val2 {
		if val2[i] != "" {
			ret := strings.Split(val2[i], "(")
			return ret[0]

		}
	}
	return val[1]
}

func findIndex(data []string, str string) int {
	for i, a := range data {
		if strings.Contains(a, str) {
			return i
		}
	}
	return 0
}

func getStringInterface(i interface{}) string {
	str := fmt.Sprintf("%v", i)
	return str
}

// PACKAGE HTTPHANDLER
//GET DOMAIN INFO
func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	ConsumeSSLApi(ctx.UserValue("server"))

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
