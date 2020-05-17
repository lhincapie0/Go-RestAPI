package dataStructure

type Domain struct {
	EndPoints []EndPoint `json:"endpoints"`
	Host      string     `json:"host"`
}

type EndPoint struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}
