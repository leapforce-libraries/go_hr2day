package hr2day

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	h_types "github.com/leapforce-libraries/go_hr2day/types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type ArbVoorwCluster struct {
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
	Hr2daySubgroepnaam string                    `json:"hr2d__Subgroepnaam__c"`
}

func (service *Service) GetArbVoorwClusters() (*[]ArbVoorwCluster, *errortools.Error) {
	s := utilities.GetTaggedTagNames("json", ArbVoorwCluster{})

	q := fmt.Sprintf("SELECT %s FROM hr2d__ArbVoorwCluster__c", s)

	var arbVoorwClusters []ArbVoorwCluster

	e := service.query(q, &arbVoorwClusters)
	if e != nil {
		return nil, e
	}

	return &arbVoorwClusters, nil
}
