package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type Token struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

func (service *Service) getToken() (*Token, *errortools.Error) {
	formData := struct {
		GrantType    string `json:"grant_type"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		"password",
		service.username,
		fmt.Sprintf("%s%s", service.password, service.securityToken),
		service.clientId,
		service.clientSecret,
	}

	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	var token Token

	t := true
	requestConfig := go_http.RequestConfig{
		Method:             http.MethodPost,
		Url:                loginUrl,
		XWwwFormUrlEncoded: &t,
		BodyModel:          formData,
		NonDefaultHeaders:  &header,
		ResponseModel:      &token,
	}

	_, _, e := service.httpService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &token, nil
}
