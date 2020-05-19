package dataStructure

//https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs-v3.md
type Domain struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Errors    []Error    `json:"errors"`
	EndPoints []EndPoint `json:"endpoints"`
}

type EndPoint struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

//{"errors":[{"field":"host","message":"qp.mandatory"}]}
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
