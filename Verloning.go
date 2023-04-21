package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Verloning struct {
	Id                   string                    `json:"Id"`
	IsDeleted            bool                      `json:"IsDeleted"`
	Name                 string                    `json:"Name"`
	CreatedDate          *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById          string                    `json:"CreatedById"`
	LastModifiedDate     *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById     string                    `json:"LastModifiedById"`
	SystemModstamp       *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate     *h_types.DateTimeMsString `json:"LastActivityDate"`
	Hr2dayEmployee       string                    `json:"hr2d__Employee__c"`
	Hr2dayArbeidsrelatie string                    `json:"hr2d__Arbeidsrelatie__c	"`
	Hr2dayArbrelVolgnr   float64                   `json:"hr2d__Arbrel_Volgnr__c"`
	Hr2dayBruto          float64                   `json:"hr2d__Bruto__c"`
	Hr2dayDeeltdAanvang  float64                   `json:"hr2d__Deeltd_Aanvang__c"`
	Hr2dayDeeltdEinde    float64                   `json:"hr2d__Deeltd_Einde__c"`
	Hr2dayGeldig         bool                      `json:"hr2d__Geldig__c"`
	Hr2dayJaarPeriode    string                    `json:"hr2d__JaarPeriode__c"`
	Hr2dayJaarIn         string                    `json:"hr2d__Jaar_in__c"`
	Hr2dayNetto          float64                   `json:"hr2d__Netto__c"`
	Hr2dayPeriodeNr      float64                   `json:"hr2d__PeriodeNr__c"`
	Hr2dayPeriodeNrIn    float64                   `json:"hr2d__PeriodeNr_in__c"`
	Hr2dayPeriode        string                    `json:"hr2d__Periode__c"`
	Hr2dayUren           float64                   `json:"hr2d__Uren__c"`
	Hr2dayUurloon        float64                   `json:"hr2d__Uurloon__c"`
}

func (service *Service) GetVerlonings() (*[]Verloning, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Verloning{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Verloning__c", s)

	var verlonings []Verloning

	e := service.query(q, &verlonings)
	if e != nil {
		return nil, e
	}

	return &verlonings, nil
}
