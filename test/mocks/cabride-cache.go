package mocks

import (
	"time"

	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabrideMockCache is an mock for pkg/rides/CabRideCache.
type CabrideMockCache struct {
	FetchCall struct {
		Receives struct {
			Medallions []string
			PickupDate time.Time
		}
		Returns struct {
			TripsData []*pb.TripData
			Error     error
		}
	}

	CacheCall struct {
		Receives struct {
			Trips      []*pb.TripData
			PickupDate time.Time
		}
		Returns struct {
			Error error
		}
	}

	ClearCall struct {
		Receives struct {
			PickupDate time.Time
		}
		Returns struct {
			Error error
		}
	}

	FlushCall struct {
		Receives struct {
		}
		Returns struct {
			Error error
		}
	}
}

// Fetch Mock.
func (cm *CabrideMockCache) Fetch(medallions []string, d time.Time) ([]*pb.TripData, error) {
	cm.FetchCall.Receives.Medallions = medallions
	cm.FetchCall.Receives.PickupDate = d
	return cm.FetchCall.Returns.TripsData, cm.FetchCall.Returns.Error
}

// Cache Mock.
func (cm *CabrideMockCache) Cache(trips []*pb.TripData, d time.Time) error {
	cm.CacheCall.Receives.Trips = trips
	cm.CacheCall.Receives.PickupDate = d
	return cm.CacheCall.Returns.Error
}

// Clear Mock.
func (cm *CabrideMockCache) Clear(d time.Time) error {
	cm.ClearCall.Receives.PickupDate = d
	return cm.ClearCall.Returns.Error
}

// Flush Mock.
func (cm *CabrideMockCache) Flush() error {
	return cm.FlushCall.Returns.Error
}
