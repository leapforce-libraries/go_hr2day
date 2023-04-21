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
	Hr2dayAanvangArbrel   *h_types.DateString       `json:"hr2d__Aanvang_arbrel__c"`
	Hr2dayArbVoorwCluster string                    `json:"hr2d__ArbVoorwCluster__c"`
	Hr2dayAuto            string                    `json:"hr2d__Auto__c"`
	Hr2dayAutoDatum       *h_types.DateString       `json:"hr2d__AutoDatum__c"`
	Hr2dayAutoKenteken    string                    `json:"hr2d__AutoKenteken__c"`
	Hr2dayAutoMerkType    string                    `json:"hr2d__AutoMerkType__c"`
	Hr2dayAutoRdnGnBijt   string                    `json:"hr2d__AutoRdnGnBijt__c"`
	Hr2dayAutoTypering    string                    `json:"hr2d__AutoTypering__c"`
	Hr2dayAutoWaarde      *float64                  `json:"hr2d__AutoWaarde__c"`
	Hr2dayContractAanvang *h_types.DateString       `json:"hr2d__Contract_Aanvang__c"`
	Hr2dayContractBepTijd string                    `json:"hr2d__Contract_bep_tijd__c"`
	Hr2dayContractEinde   *h_types.DateString       `json:"hr2d__Contract_Einde__c"`
	Hr2dayContractVolgnr  *float64                  `json:"hr2d__Contract_Volgnr__c"`
	Hr2dayCostCenter      string                    `json:"hr2d__CostCenter__c"`
	Hr2dayDeeltijdFactor  *float64                  `json:"hr2d__DeeltijdFactor__c"`
	Hr2dayDepartment      string                    `json:"hr2d__Department__c"`
	Hr2dayEindeArbrel     *h_types.DateString       `json:"hr2d__Einde_arbrel__c"`
	Hr2dayEmployee        string                    `json:"hr2d__Employee__c"`
	Hr2dayFiets           string                    `json:"hr2d__Fiets__c"`
	Hr2dayFietsWaarde     *float64                  `json:"hr2d__FietsWaarde__c"`
	Hr2dayFunctie         string                    `json:"hr2d__Functie__c"`
	Hr2dayGeldigTot       *h_types.DateString       `json:"hr2d__Geldig_tot__c"`
	Hr2dayGeldigVan       *h_types.DateString       `json:"hr2d__Geldig_van__c"`
	Hr2dayJob             string                    `json:"hr2d__Job__c"`
	Hr2dayLocation        string                    `json:"hr2d__Location__c"`
	Hr2daySalaris         *float64                  `json:"hr2d__Salaris__c"`
	Hr2daySchaal          string                    `json:"hr2d__Schaal__c"`
	Hr2dayUrenWeek        *float64                  `json:"hr2d__UrenWeek__c"`
	Hr2dayVolgnummer      *float64                  `json:"hr2d__Volgnummer__c"`
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
