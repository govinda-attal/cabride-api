package mocks

import (
	"time"

	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabrideMockDS is an mock pkg/rides/CabRide datastore.
type CabrideMockDS struct {
	FetchCall struct {
		Receives struct {
			Medallions []string
			PickupDate time.Time
		}
		Returns struct {
			Trips []*pb.TripData
			Error error
		}
	}

	CountCall struct {
		Receives struct {
			Medallions []string
			PickupDate time.Time
		}
		Returns struct {
			TotalTripCount int
			Error          error
		}
	}
}

// Fetch Mock.
func (cm *CabrideMockDS) Fetch(medallions []string, d time.Time) ([]*pb.TripData, error) {
	cm.FetchCall.Receives.Medallions = medallions
	cm.FetchCall.Receives.PickupDate = d
	return cm.FetchCall.Returns.Trips, cm.FetchCall.Returns.Error
}

// Count Mock.
func (cm *CabrideMockDS) Count(medallions []string, d time.Time) (int, error) {
	cm.CountCall.Receives.Medallions = medallions
	cm.CountCall.Receives.PickupDate = d
	return cm.CountCall.Returns.TotalTripCount, cm.CountCall.Returns.Error
}
