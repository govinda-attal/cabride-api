package rides

import (
	"time"

	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabRide serves the domain service.
type CabRide interface {
	// Fetch return trip data information for given set of medallions on a given pickup date.
	Fetch(medallions []string, d time.Time) ([]*pb.TripData, error)
	// Count return a total count of trips for given set of medallions on a given pickup date.
	Count(medallions []string, d time.Time) (int, error)
}

// CabRideCache serves the domain service.
type CabRideCache interface {
	// Fetch return trip data information for given set of medallions on a given pickup date.
	// It also returns a list of medallions that were not found in the cache.
	Fetch(medallions []string, d time.Time) ([]*pb.TripData, []string, error)
	// Cache will cache the trip data set for given date.
	Cache(trips []*pb.TripData, d time.Time) error
	// Clear removes entries for a given date
	Clear(d time.Time) error
	// Flush resets the cache
	Flush() error
}
