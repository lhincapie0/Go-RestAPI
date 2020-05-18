package infoHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"

	"github.com/badoux/goscraper"
	ds "github.com/lhincapie0/go-restAPI/API/dataStructure"
	"github.com/likexian/whois-go"
	"github.com/valyala/fasthttp"
)

var domain ds.Domain
var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

const serverInfoPath string = "https://api.ssllabs.com/api/v3/analyze?host="
const COUNTRY string = "Country"
const ORGANIZATION string = "Organization"

func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	fmt.Println("SENDING INFO")
	ConsumeSSLApi(ctx.UserValue("server"))
	resp := BuildDomainResponse()
	obj, err := json.Marshal(resp)
	var obj2 string
	if err != nil {
		json.Unmarshal([]byte(obj), &obj2)
		fmt.Fprintf(ctx, obj2)

	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	if err := json.NewEncoder(ctx).Encode(resp); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

//PACKAGE GET DATA
func ConsumeSSLApi(server interface{}) {
	serverName := getStringInterface(server)
	//path := serverInfoPath + serverName
	response, err := http.Get(serverInfoPath + serverName)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		//formatting data into domain variable
		json.Unmarshal([]byte(data), &domain)
		//Creating servers with their information

	}
}

func BuildDomainResponse() ds.DomainInfo {
	servers := HandleServerInfo()
	logo, title := GetHtmlInfo(domain.Host)
	domainResponse := ds.DomainInfo{
		Servers:  servers,
		SslGrade: calculateSslGrade(servers),
		Logo:     logo,
		Title:    title,
	}
	return domainResponse
}

func HandleServerInfo() []ds.Server {

	//CREATE SERVER ARRAY
	var servers []ds.Server
	for i := range domain.EndPoints {
		country, owner := WhoIs(domain.EndPoints[i].IpAddress)
		s := ds.Server{
			IpAddress: domain.EndPoints[i].IpAddress,
			Grade:     domain.EndPoints[i].Grade,
			Country:   country,
			Owner:     owner,
		}
		servers = append(servers, s)
	}
	return servers
}

//CODE COPIED FROM THE  github.com/badoux/goscraper EXAMPLE
func GetHtmlInfo(url string) (string, string) {
	url = "http://" + url
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		fmt.Println(err)
	}
	return s.Preview.Icon, s.Preview.Title
}

func WhoIs(ip string) (string, string) {
	result, err := whois.Whois(ip)
	if err == nil {
		data := strings.Split(result, "\n")

		indexCountry := findIndex(data, COUNTRY)
		country := findValueWhoIs(data, indexCountry)
		indexOrganization := findIndex(data, ORGANIZATION)
		organization := findValueWhoIs(data, indexOrganization)
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

func calculateSslGrade(servers []ds.Server) string {
	def := "A"

	for i := range servers {

		if servers[i].Grade > def {
			def = servers[i].Grade
		}
		if strings.Contains(servers[i].Grade, def) {
			def = servers[i].Grade
		}
	}

	return def
}
