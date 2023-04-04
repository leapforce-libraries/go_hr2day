package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type SickLeave struct {
	Id                        string                    `json:"Id"`
	IsDeleted                 bool                      `json:"IsDeleted"`
	Name                      string                    `json:"Name"`
	CreatedDate               *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById               string                    `json:"CreatedById"`
	LastModifiedDate          *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById          string                    `json:"LastModifiedById"`
	SystemModstamp            *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate          *h_types.DateString       `json:"LastActivityDate"`
	Hr2dEmployeeC             string                    `json:"hr2d__Employee__c"`
	Hr2dCaseManagerC          string                    `json:"hr2d__CaseManager__c"`
	Hr2dEndDateC              *h_types.DateString       `json:"hr2d__EndDate__c"`
	Hr2dSickClassificationC   string                    `json:"hr2d__SickClassification__c"`
	Hr2dStartDateC            *h_types.DateString       `json:"hr2d__StartDate__c"`
	Hr2dArbrelVolgnrC         *float64                  `json:"hr2d__ArbrelVolgnr__c"`
	Hr2dDateRelatedCaseC      *h_types.DateString       `json:"hr2d__DateRelatedCase__c"`
	Hr2dDepartmentC           string                    `json:"hr2d__Department__c"`
	Hr2dClassificationPickedC string                    `json:"hr2d__ClassificationPicked__c"`
	Hr2dHoursFirstDayC        *float64                  `json:"hr2d__HoursFirstDay__c"`
	Hr2dNursingAddressC       bool                      `json:"hr2d__NursingAddress__c"`
	Hr2dActualDueDateC        *h_types.DateString       `json:"hr2d__ActualDueDate__c"`
	Hr2dExpectedDueDateC      *h_types.DateString       `json:"hr2d__ExpectedDueDate__c"`
}

func (service *Service) GetSickLeaves() (*[]SickLeave, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", SickLeave{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__SickLeave__c", s)

	var sickLeaves []SickLeave

	e := service.query(q, &sickLeaves)
	if e != nil {
		return nil, e
	}

	return &sickLeaves, nil
}
