package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Employer struct {
	Id                 string                    `json:"Id"`
	OwnerId            string                    `json:"OwnerId"`
	IsDeleted          bool                      `json:"IsDeleted"`
	Name               string                    `json:"Name"`
	CreatedDate        *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById        string                    `json:"CreatedById"`
	LastModifiedDate   *h_types.DateTimeMsString `json:"LastModifiedDate"`
	SystemModstamp     *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate   *h_types.DateTimeMsString `json:"LastActivityDate"`
	LastViewedDate     *h_types.DateTimeMsString `json:"LastViewedDate"`
	LastReferencedDate *h_types.DateTimeMsString `json:"LastReferencedDate"`
	Hr2dayFullName     string                    `json:"hr2d__FullName__c"`
}

func (service *Service) GetEmployers() (*[]Employer, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Employer{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Employer__c", s)

	var employers []Employer

	e := service.query(q, &employers)
	if e != nil {
		return nil, e
	}

	return &employers, nil
}
