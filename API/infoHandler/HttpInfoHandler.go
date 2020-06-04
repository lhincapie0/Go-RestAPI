package infohandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/badoux/goscraper"
	"github.com/lhincapie0/Go-RestAPI/API/database"

	ds "github.com/lhincapie0/go-restAPI/API/dataStructure"
	"github.com/likexian/whois-go"
	"github.com/valyala/fasthttp"
)

var domain ds.Domain
var domainCheck ds.Domain

var connectionDB *sql.DB

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)
var fromCache bool

const serverInfoPath string = "https://api.ssllabs.com/api/v3/analyze?host="
const param1 string = "&onCache=on&max-age=1"
const param2 string = "&startNew=on"

//COUNTRY ... for who is - country
const COUNTRY string = "Country"

//ORGANIZATION ... for who is - organization
const ORGANIZATION string = "Organization"

//READY ... state where the api result is ready to use
const READY string = "READY"

//INVALID ... error if the input domain is invalidad
const INVALID string = "Invalid domain"

//FULL ... error if the server is busy
const FULL string = "Running at full capacity. Please try again later."

//HTTPInfoHandler ... Instantiates the database
func HTTPInfoHandler(db *sql.DB) {
	connectionDB = db
}

//GetDomainInfo ... main method to build the endpoint respond.
func GetDomainInfo(ctx *fasthttp.RequestCtx) {
	ConsumeSSLApi(ctx.UserValue("server"), param1)
	dom := getStringInterface(ctx.UserValue("server"))
	fmt.Println(dom)
	if len(domain.Errors) > 0 || domain.Status == "ERROR" {
		fmt.Println("ERRORS  ", domain.Errors)
		if domain.Status == "ERROR" {
			json.NewEncoder(ctx).Encode(INVALID)
		} else if domain.Errors[0].Message == FULL {
			json.NewEncoder(ctx).Encode(FULL)
		} else {
			json.NewEncoder(ctx).Encode("Unknown error")
		}
	} else {

		resp := BuildDomainResponse(dom)
		///	database.AddDomain(connectionDB, getStringInterface(ctx.UserValue("server")), resp.SslGrade)
		database.AddDomainInfo(connectionDB, getStringInterface(ctx.UserValue("server")), resp)
		ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
		fmt.Println("Information sent")
		if err := json.NewEncoder(ctx).Encode(resp); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}

}

func checkIsDown(dom string) bool {
	var st string = "http://www." + dom
	response, _ := http.Get(st)
	fmt.Println(st)
	fmt.Println(response)

	if response != nil {
		return false
	} else {
		return true
	}
}

func ConsumeSSLApi(server interface{}, params string) {
	serverName := getStringInterface(server)
	response, err := http.Get(serverInfoPath + serverName + params)
	fromCache = false
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		fromCache = true
		//formatting data into domain variable
		json.Unmarshal([]byte(data), &domain)
		fromCache = true

		for domain.Status != "READY" {
			if domain.Status == "ERROR" {
				fmt.Println("error..............")
				break
			} else {
				response, err := http.Get(serverInfoPath + serverName)

				if err != nil {
					fmt.Printf("The HTTP request failed with error %s\n", err)
				} else {
					fromCache = false

					data, _ := ioutil.ReadAll(response.Body)
					//formatting data into domain variable
					json.Unmarshal([]byte(data), &domain)

				}
			}

		}
	}
}

func BuildDomainResponse(dom string) ds.DomainInfo {
	fmt.Println(dom)
	servers := HandleServerInfo()
	logo, title := GetHtmlInfo(domain.Host)
	fmt.Println("BUILDING DOMAAIN")
	var isDown bool
	var changed bool
	var previousSsl string
	var actualSslGrade string
	actualSslGrade = calculateSslGrade(servers)
	isDown = checkIsDown(dom)
	if fromCache {
		previousSsl, changed = evaluateChanges(domain.Host, servers, actualSslGrade)
	} else {
		changed = false
		previousSsl = actualSslGrade
	}

	domainResponse := ds.DomainInfo{
		Servers:          servers,
		SeversChanged:    changed,
		SslGrade:         actualSslGrade,
		PreviousSslGrade: previousSsl,
		Logo:             logo,
		Title:            title,
		IsDown:           isDown,
	}
	return domainResponse
}

func evaluateChanges(host string, previousServers []ds.Server, actualSslGrade string) (string, bool) {
	var changes bool = false
	//domain search for one hour or less saved in domainCheck
	domainCheck = domain
	fmt.Println("EVALUATING CHANGES IN DOMAAIN")

	//actual (now) search saved in domain
	//Param2 indicates that the result has to be a new one and not the one in the cache
	ConsumeSSLApi(host, param2)
	servers := HandleServerInfo()
	var previousSsl string
	if len(servers) != len(previousServers) {
		changes = true
	} else {
		for _, s := range servers {
			if !ServerExists(previousServers, s) {
				changes = true
			}
		}
	}
	if changes {
		previousSsl = calculateSslGrade(servers)
	} else {
		previousSsl = actualSslGrade
	}
	return previousSsl, changes
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
	var icon string
	var title string
	s, err := goscraper.Scrape(url, 5)
	if err == nil {

	}
	if s.Preview.Icon == "" {
		icon = "No icon"
	} else {
		icon = s.Preview.Icon
	}

	if s.Preview.Title == "" {
		title = "No icon"
	} else {
		title = s.Preview.Title
	}
	return icon, title
}

func WhoIs(ip string) (string, string) {
	result, err := whois.Whois(ip)
	if err == nil {
		indexCountry := findIndex(result, COUNTRY)

		country := findValueWhoIs(result, indexCountry)
		indexOrganization := findIndex(result, ORGANIZATION)
		organization := findValueWhoIs(result, indexOrganization)
		return country, organization

	}
	return "", ""
}

func findValueWhoIs(result string, index int) string {
	if index == 0 {
		return "null"
	} else {
		data := strings.Split(result, "\n")

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

}

func findIndex(result string, str string) int {
	data := strings.Split(result, "\n")
	if strings.Contains(strings.ToUpper(result), strings.ToUpper(str)) {

		for i, a := range data {
			if strings.Contains(strings.ToUpper(a), strings.ToUpper(str)) {
				return i

			}
		}
	}
	return 0

}

func GetDomainsHistory(ctx *fasthttp.RequestCtx) {
	var hosts []string
	hosts = database.GetDomains(connectionDB)
	repo := ds.HostRepo{
		Items: hosts,
	}

	obj, err := json.Marshal(repo)
	var obj2 string
	if err != nil {
		json.Unmarshal([]byte(obj), &obj2)
		fmt.Fprintf(ctx, obj2)
	}
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	if err := json.NewEncoder(ctx).Encode(repo); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

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

func ServerExists(servers []ds.Server, item ds.Server) bool {
	for _, s := range servers {
		if s == item {
			return true
		}
	}
	return false

}
