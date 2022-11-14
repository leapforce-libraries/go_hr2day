package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type LooncompOutput struct {
	Id                  string                    `json:"Id"`
	IsDeleted           bool                      `json:"IsDeleted"`
	Name                string                    `json:"Name"`
	CreatedDate         *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById         string                    `json:"CreatedById"`
	LastModifiedDate    *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById    string                    `json:"LastModifiedById"`
	SystemModstamp      *h_types.DateTimeMsString `json:"SystemModstamp"`
	Hr2dayVerloning     string                    `json:"hr2d__Verloning__c"`
	Hr2dayAantal        *float64                  `json:"hr2d__Aantal__c"`
	Hr2dayBedrag        *float64                  `json:"hr2d__Bedrag__c"`
	Hr2dayFactor        *float64                  `json:"hr2d__Factor__c"`
	Hr2dayLooncomponent string                    `json:"hr2d__Looncomponent__c"`
	Hr2dayTarief        *float64                  `json:"hr2d__Tarief__c"`
}

func (service *Service) GetLooncompOutputs() (*[]LooncompOutput, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", LooncompOutput{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__LooncompOutput__c", s)

	var looncompOutputs []LooncompOutput

	e := service.query(q, &looncompOutputs)
	if e != nil {
		return nil, e
	}

	return &looncompOutputs, nil
}
