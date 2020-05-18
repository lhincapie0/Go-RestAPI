package dataStructure

type DomainInfo struct {
	Servers []Server `json:"servers"`
	//	SeversChanged   bool     `json:"servers_changed"`
	SslGrade string `json:"ssl_grade"`
	//	PreviousSslGrade bool   `json:"previous_ssl_grade"`
	Logo  string `json:"logo"`
	Title string `json:"title"`
	//	IsDown bool   `json:"is_down"`
}

type Server struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}
