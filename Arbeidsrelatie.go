package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Arbeidsrelatie struct {
	Id                    string                    `json:"Id"`
	IsDeleted             bool                      `json:"IsDeleted"`
	Name                  string                    `json:"Name"`
	CreatedDate           *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById           string                    `json:"CreatedById"`
	LastModifiedDate      *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById      string                    `json:"LastModifiedById"`
	SystemModstamp        *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate      *h_types.DateTimeMsString `json:"LastActivityDate"`
	Hr2dayEmployee        string                    `json:"hr2d__Employee__c"`
	Hr2dayAanvangArbrel   *h_types.DateString       `json:"hr2d__Aanvang_arbrel__c"`
	Hr2dayArbVoorwCluster string                    `json:"hr2d__ArbVoorwCluster__c"`
	Hr2dayCostCenter      string                    `json:"hr2d__CostCenter__c"`
	Hr2dayDepartment      string                    `json:"hr2d__Department__c"`
	Hr2dayEindeArbrel     *h_types.DateString       `json:"hr2d__Einde_arbrel__c	"`
	Hr2dayVolgnummer      float64                   `json:"hr2d__Volgnummer__c"`
	Hr2dayJob             string                    `json:"hr2d__Job__c"`
	Hr2dayDeeltijdFactor  float64                   `json:"hr2d__DeeltijdFactor__c"`
	Hr2dayLocation        string                    `json:"hr2d__Location__c"`
}

func (service *Service) GetArbeidsrelaties() (*[]Arbeidsrelatie, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Arbeidsrelatie{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Arbeidsrelatie__c", s)

	var arbeidsrelaties []Arbeidsrelatie

	e := service.query(q, &arbeidsrelaties)
	if e != nil {
		return nil, e
	}

	return &arbeidsrelaties, nil
}
