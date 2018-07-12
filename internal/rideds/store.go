package rideds

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jmoiron/sqlx"

	"github.com/govinda-attal/cabride-api/internal/provider"

	"github.com/govinda-attal/cabride-api/pkg/core/status"
	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabRideStore is an implementation of pkg/ride/CabRide.
type CabRideStore struct {
	DB *sqlx.DB
}

// NewCabRideStore returns new instance of datastore service.
func NewCabRideStore() *CabRideStore {
	return &CabRideStore{provider.DB()}
}

// Fetch return trip data information for given set of medallions on a given pickup date.
func (cr *CabRideStore) Fetch(medallions []string, d time.Time) ([]*pb.TripData, error) {
	// Getting count of recrds if they exist for search criteria is helpful to decide size of the slice of TripData.
	// To my knowledge database/sql(x) don't provide a clean way to return count of rows of the result set returned, as it is a cursor.
	// There can be a fair assumption for a city like new york there can be large number of cab rides for a given date
	// So pre-allocating the slize of TripData struct will be more efficient rather than appending the slice
	count, err := cr.Count(medallions, d)
	if err != nil {
		return nil, status.ErrInternal.WithMessage(err.Error())
	}
	if count == 0 {
		return nil, status.ErrNotFound
	}

	db := cr.DB
	qryStmt := `SELECT medallion, hack_license, vendor_id, rate_code, store_and_fwd_flag, 
					pickup_datetime, dropoff_datetime, passenger_count, trip_time_in_secs, trip_distance FROM cab_trip_data
					WHERE medallion IN (:medallions) AND DATE(pickup_datetime) = :pickupDate ORDER BY medallion;`

	arg := map[string]interface{}{
		"pickupDate": time.Time(d).Format("2006-01-02"),
		"medallions": medallions,
	}

	query, args, err := sqlx.Named(qryStmt, arg)
	query, args, err = sqlx.In(query, args...)
	query = db.Rebind(query)
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, status.ErrInternal.WithMessage(err.Error())
	}

	// Allocating a slice of trips with right size.
	trips := make([]*pb.TripData, count)
	i := 0

	for rows.Next() {
		var trip pb.TripData
		var pt, dt time.Time
		err = rows.Scan(&trip.Medallion, &trip.HackLicense,
			&trip.VendorID, &trip.RateCode, &trip.StoreFwdFlag,
			&pt, &dt, &trip.PassengerCount, &trip.TripTime, &trip.TripDistance)

		if err != nil {
			return nil, status.ErrInternal.WithMessage(err.Error())
		}
		if i >= count { // In  an unlikely scenario of fraction of second query with actual data-set resulted in larger number than initial  reported count.
			newSlice := make([]*pb.TripData, i+1)
			copy(newSlice, trips)
			trips = newSlice
		}
		// There is no clean conversion from time.Time returned by DB Query to protobuf timestamp.Timestamp.
		trip.Pickup, err = ptypes.TimestampProto(pt)
		trip.Dropoff, err = ptypes.TimestampProto(dt)
		if err != nil {
			return nil, status.ErrInternal.WithMessage(err.Error())
		}
		trips[i] = &trip
		i = i + 1
	}
	rows.Close()
	return trips, nil
}

// Count return a total count of trips for given set of medallions on a given pickup date.
func (cr *CabRideStore) Count(medallions []string, d time.Time) (int, error) {
	count := 0
	db := cr.DB
	countStmt := `SELECT COUNT(*) FROM cab_trip_data
					WHERE medallion IN (:medallions) AND DATE(pickup_datetime) = :pickupDate;`
	arg := map[string]interface{}{
		"pickupDate": time.Time(d).Format("2006-01-02"),
		"medallions": medallions,
	}

	query, args, err := sqlx.Named(countStmt, arg)
	query, args, err = sqlx.In(query, args...)
	query = db.Rebind(query)
	err = db.Get(&count, query, args...)

	if err != nil {
		return 0, status.ErrInternal.WithMessage(err.Error())
	}
	return count, nil
}
