package dataStructure

type DomainHistory struct {
	Items []DomainHistoryElement `json:"items"`
}

type DomainHistoryElement struct {
	Host string        `json:"host"`
	Info DomainInfoStr `json:"info"`
}

type DomainInfoStr struct {
	Servers          string `json:"servers"`
	SeversChanged    string `json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:"logo"`
	Title            string `json:"title"`
	IsDown           string `json:"is_down"`
}
type DomainInfo struct {
	Servers          []Server `json:"servers"`
	SeversChanged    bool     `json:"servers_changed"`
	SslGrade         string   `json:"ssl_grade"`
	PreviousSslGrade string   `json:"previous_ssl_grade"`
	Logo             string   `json:"logo"`
	Title            string   `json:"title"`
	IsDown           bool     `json:"is_down"`
}

type Server struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}

type HostRepo struct {
	Items []string `json:"items"`
}
