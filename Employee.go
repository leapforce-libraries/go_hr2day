package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Employee struct {
	Id                    string                    `json:"Id"`
	OwnerId               string                    `json:"OwnerId"`
	IsDeleted             bool                      `json:"IsDeleted"`
	Name                  string                    `json:"Name"`
	CreatedDate           *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById           string                    `json:"CreatedById"`
	LastModifiedDate      *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById      string                    `json:"LastModifiedById"`
	SystemModstamp        *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate      *h_types.DateTimeMsString `json:"LastActivityDate"`
	LastViewedDate        *h_types.DateTimeMsString `json:"LastViewedDate"`
	LastReferencedDate    *h_types.DateTimeMsString `json:"LastReferencedDate"`
	Hr2dayASurname        string                    `json:"hr2d__A_Surname__c"`
	Hr2dayAName           string                    `json:"hr2d__A_name__c"`
	Hr2dayEmplNr          float64                   `json:"hr2d__EmplNr__c"`
	Hr2dayEmployer        string                    `json:"hr2d__Employer__c"`
	Hr2dayFirstName       string                    `json:"hr2d__FirstName__c"`
	Hr2dayHireDate        *h_types.DateString       `json:"hr2d__HireDate__c"`
	Hr2dayInitials        string                    `json:"hr2d__Initials__c"`
	Hr2dayNickname        string                    `json:"hr2d__Nickname__c"`
	Hr2dayPhone2          string                    `json:"hr2d__Phone2__c"`
	Hr2dayPhone           string                    `json:"hr2d__Phone__c"`
	Hr2dayPrefix          string                    `json:"hr2d__Prefix__c"`
	Hr2daySurname         string                    `json:"hr2d__Surname__c"`
	Hr2dayTerminationDate *h_types.DateString       `json:"hr2d__TerminationDate__c"`
	Hr2dayEmailWork       string                    `json:"hr2d__EmailWork__c"`
	Hr2dayDepartmentToday string                    `json:"hr2d__DepartmentToday__c"`
	Hr2dayJobToday        string                    `json:"hr2d__JobToday__c"`
	Hr2dayPhone3          string                    `json:"hr2d__Phone3__c"`
}

func (service *Service) GetEmployees() (*[]Employee, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Employee{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Employee__c", s)

	var employees []Employee

	e := service.query(q, &employees)
	if e != nil {
		return nil, e
	}

	return &employees, nil
}
