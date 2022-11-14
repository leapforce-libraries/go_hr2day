package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type LooncompDefinitie struct {
	Id                 string                    `json:"Id"`
	OwnerId            string                    `json:"OwnerId"`
	IsDeleted          bool                      `json:"IsDeleted"`
	Name               string                    `json:"Name"`
	CreatedDate        *h_types.DateTimeMsString `json:"CreatedDate"`
	CreatedById        string                    `json:"CreatedById"`
	LastModifiedDate   *h_types.DateTimeMsString `json:"LastModifiedDate"`
	LastModifiedById   string                    `json:"LastModifiedById"`
	SystemModstamp     *h_types.DateTimeMsString `json:"SystemModstamp"`
	LastViewedDate     *h_types.DateTimeMsString `json:"LastViewedDate"`
	LastReferencedDate *h_types.DateTimeMsString `json:"LastReferencedDate"`
}

func (service *Service) GetLooncompDefinities() (*[]LooncompDefinitie, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", LooncompDefinitie{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__LooncompDefinitie__c", s)

	var employees []LooncompDefinitie

	e := service.query(q, &employees)
	if e != nil {
		return nil, e
	}

	return &employees, nil
}
