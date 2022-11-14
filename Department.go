package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Department struct {
	Id               string                    `json:"Id"`
	IsDeleted        bool                      `json:"IsDeleted"`
	Name             string                    `json:"Name"`
	CreatedDate      *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById      string                    `json:"CreatedById"`
	LastModifiedDate *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById string                    `json:"LastModifiedById"`
	SystemModstamp   *h_types.DateTimeMsString `json:"SystemModstamp"`
	Hr2dayEmployer   string                    `json:"hr2d__Employer__c"`
}

func (service *Service) GetDepartments() (*[]Department, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Department{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Department__c", s)

	var departments []Department

	e := service.query(q, &departments)
	if e != nil {
		return nil, e
	}

	return &departments, nil
}
