package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

const (
	apiName  string = "HR2day"
	loginUrl string = "https://login.salesforce.com/services/oauth2/token"
	queryUrl string = "https://%s.cloudforce.com/services/data/v56.0/query"
)

type Service struct {
	domain        string
	username      string
	password      string
	securityToken string
	clientId      string
	clientSecret  string
	httpService   *go_http.Service
	token         *Token
}

type ServiceConfig struct {
	Domain        string
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
		domain:        serviceConfig.Domain,
		username:      serviceConfig.Username,
		password:      serviceConfig.Password,
		securityToken: serviceConfig.SecurityToken,
		clientId:      serviceConfig.ClientId,
		clientSecret:  serviceConfig.ClientSecret,
		httpService:   httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	if service.token == nil {
		token, e := service.getToken()
		if e != nil {
			return nil, nil, e
		}

		service.token = token
	}

	// add access token to header
	if requestConfig.NonDefaultHeaders == nil {
		requestConfig.NonDefaultHeaders = &http.Header{}
	}
	requestConfig.NonDefaultHeaders.Set("Authorization", fmt.Sprintf("Bearer %s", service.token.AccessToken))

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
