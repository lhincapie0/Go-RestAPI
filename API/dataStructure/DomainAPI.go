package dataStructure

type DomainInfo struct {
	Servers  []Server `json:"servers"`
	SslGrade string   `json:"ssl_grade"`
	Logo     string   `json:"logo"`
	Title    string   `json:"title"`
}

type Server struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}
