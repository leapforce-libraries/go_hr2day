package hr2day

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName string = "HR2day"
	apiURL  string = "https://api.neverbounce.com/v4"
)

type Service struct {
	username      string
	password      string
	securityToken string
	clientId      string
	clientSecret  string
	httpService   *go_http.Service
}

type ServiceConfig struct {
	Username      string
	Password      string
	SecurityToken string
	ClientId      string
	ClientSecret  string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		username:      serviceConfig.Username,
		password:      serviceConfig.Password,
		securityToken: serviceConfig.SecurityToken,
		clientId:      serviceConfig.ClientId,
		clientSecret:  serviceConfig.ClientSecret,
		httpService:   httpService,
	}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add Api key
	_url, err := url.Parse(requestConfig.Url)
	if err != nil {
		return nil, nil, errortools.ErrorMessage(err)
	}
	query := _url.Query()
	//query.Set("key", service.clientId)

	(*requestConfig).Url = fmt.Sprintf("%s://%s%s?%s", _url.Scheme, _url.Host, _url.Path, query.Encode())

	return service.httpService.HttpRequest(requestConfig)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.clientId
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
