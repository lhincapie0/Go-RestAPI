package servers

import (
	"fmt"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

const serverInfoPath string = "https://api.ssllabs.com/api/v3/analyze?host="

func whoIsInfo(server string) string {

	result, err := whois.Whois(server)
	if err == nil {
		fmt.Println(result)

	}
	return result
}

func whoIsInfoParse(server string) {
	info := whoIsInfo(server)
	result, err := whoisparser.Parse(info)
	if err == nil {
		fmt.Println(result.Administrative.Country)
		fmt.Println(result.Registrar.Organization)

	} else {
		fmt.Println("hola")
	}
}

func try() {
	fmt.Println("can import now")
}
