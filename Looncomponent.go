package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Looncomponent struct {
	Id                   string                    `json:"Id"`
	IsDeleted            bool                      `json:"IsDeleted"`
	Name                 string                    `json:"Name"`
	CreatedDate          *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById          string                    `json:"CreatedById"`
	LastModifiedDate     *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById     string                    `json:"LastModifiedById"`
	SystemModstamp       *h_types.DateTimeMsString `json:"SystemModstamp"`
	Hr2dayArbeidsrelatie string                    `json:"hr2d__Arbeidsrelatie__c"`
	Hr2dayAantal         *float64                  `json:"hr2d__Aantal__c"`
	Hr2dayBedrag         *float64                  `json:"hr2d__Bedrag__c"`
	Hr2dayFactor         *float64                  `json:"hr2d__Factor__c"`
	Hr2dayGeldigTot      *h_types.DateString       `json:"hr2d__Geldig_tot__c"`
	Hr2dayGeldigVan      *h_types.DateString       `json:"hr2d__Geldig_van__c"`
	Hr2dayLooncomponent  string                    `json:"hr2d__Looncomponent__c"`
	Hr2dayTarief         *float64                  `json:"hr2d__Tarief__c"`
	Hr2dayCostCenter     string                    `json:"hr2d__CostCenter__c"`
	Hr2dayDepartment     string                    `json:"hr2d__Department__c"`
}

func (service *Service) GetLooncomponents() (*[]Looncomponent, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Looncomponent{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Looncomponent__c", s)

	var looncomponents []Looncomponent

	e := service.query(q, &looncomponents)
	if e != nil {
		return nil, e
	}

	return &looncomponents, nil
}
