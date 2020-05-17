package datamodel

import (
	"model.go"
)

func (d *Domain) OK() error {
	if len(d.Title) == 0 {
		return missingFieldError("title")
	}

	if len(d.Logo) == 0 {
		return missingFieldError("logo")
	}

	if len(d.SslGrade) == 0 {
		return missingFieldError("ssl_grade")
	}

	if len(d.IsDown) == nil {
		return missingFieldError("is_down")
	}
	
	if len(d.PreviousSslGrade) == 0 {
		return missingFieldError("previous_ssl_grade")
	}
	
	if len(d.ServersChanged) == nil {
		return missingFieldError("servers_changed")
	}
	return nil
}

