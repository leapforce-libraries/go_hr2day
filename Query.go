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
	TotalSize      int64             `json:"totalSize"`
	Done           bool              `json:"done"`
	NextRecordsUrl string            `json:"nextRecordsUrl"`
	Records        []json.RawMessage `json:"records"`
}

func (service *Service) query(query string, model interface{}) *errortools.Error {
	values := url.Values{}
	values.Set("q", query)

	url := fmt.Sprintf("%s%s?%s", fmt.Sprintf(baseUrl, service.domain), queryPath, values.Encode())

	var records []json.RawMessage

	for {
		var qr queryResponse
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &qr,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return e
		}

		records = append(records, qr.Records...)

		if qr.NextRecordsUrl == "" {
			break
		}

		url = fmt.Sprintf("%s%s", fmt.Sprintf(baseUrl, service.domain), qr.NextRecordsUrl)
	}

	b, err := json.Marshal(records)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	err = json.Unmarshal(b, model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}
