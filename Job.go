package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Job struct {
	Id               string                    `json:"Id"`
	IsDeleted        bool                      `json:"IsDeleted"`
	Name             string                    `json:"Name"`
	CreatedDate      *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById      string                    `json:"CreatedById"`
	LastModifiedDate *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById string                    `json:"LastModifiedById"`
	SystemModstamp   *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastActivityDate *h_types.DateTimeMsString `json:"LastActivityDate"`
	Hr2dayEmployer   string                    `json:"hr2d__Employer__c"`
}

func (service *Service) GetJobs() (*[]Job, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Job{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Job__c", s)

	var jobs []Job

	e := service.query(q, &jobs)
	if e != nil {
		return nil, e
	}

	return &jobs, nil
}
