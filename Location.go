package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type Location struct {
	Id               string                    `json:"Id"`
	IsDeleted        bool                      `json:"IsDeleted"`
	Name             string                    `json:"Name"`
	CreatedDate      *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById      string                    `json:"CreatedById"`
	LastModifiedDate *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById string                    `json:"LastModifiedById"`
	SystemModstamp   *h_types.DateTimeMsString `json:"SystemModstamp"`
}

func (service *Service) GetLocations() (*[]Location, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", Location{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__Location__c", s)

	var locations []Location

	e := service.query(q, &locations)
	if e != nil {
		return nil, e
	}

	return &locations, nil
}
