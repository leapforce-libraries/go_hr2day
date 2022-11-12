package hr2day

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type queryResponse struct {
	TotalSize int64           `json:"totalSize"`
	Done      bool            `json:"done"`
	Records   json.RawMessage `json:"records"`
}

func (service *Service) query(query string, model interface{}) *errortools.Error {
	values := url.Values{}
	values.Set("q", query)

	var qr queryResponse
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           fmt.Sprintf("%s?%s", fmt.Sprintf(queryUrl, service.domain), values.Encode()),
		ResponseModel: &qr,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	err := json.Unmarshal(qr.Records, model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}
