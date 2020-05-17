package datamodel

import (
	"fmt"
	"encoding/json"
)

var DomainStore = make(map[string]Domain)
type missingFieldError string


type Domain struct {
	Servers  []Server 
	ServersChanged bool `json:"servers_changed"`
	SslGrade string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
}

type Server struct{
	Address string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}