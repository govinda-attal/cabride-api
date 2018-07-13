// Package handler includes implementation of GRPC Service.
// This implementation relies on underlying data-store or cache services implementation pkg/rides/CabRide and pkg/rides/CabRideCache interfaces.
package handler

import (
	"log"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/govinda-attal/cabride-api/pkg/core/status"
	"github.com/govinda-attal/cabride-api/pkg/rides"
	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabTrips realizes grpc service to fetch cab trips information.
type CabTrips struct {
	crSrv   rides.CabRide
	crCache rides.CabRideCache
}

// NewCabTripsHandler returns instance of the grpc Service Handler.
func NewCabTripsHandler(cr rides.CabRide, crCache rides.CabRideCache) *CabTrips {
	return &CabTrips{crSrv: cr, crCache: crCache}
}

// Fetch returns cab trip data for a given set of medallions on a given pickup date.
func (ct *CabTrips) Fetch(ctx context.Context, rq *pb.FetchRq) (*pb.FetchRs, error) {

	err := ct.Validte(rq)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	// Validate Step will validate date
	pickup, err := time.Parse("2006-01-02", rq.Pickup)
	notFound := rq.Medallions
	tt := []*pb.TripData{}

	if !rq.NoCache {
		tt, notFound, err = ct.crCache.Fetch(rq.Medallions, pickup)
	}

	var trips []*pb.TripData

	switch {
	case len(notFound) == len(rq.Medallions):
		trips, err = ct.crSrv.Fetch(rq.Medallions, pickup)
	case len(notFound) == 0:
	case len(notFound) < len(rq.Medallions):
		trips, err = ct.crSrv.Fetch(notFound, pickup)
	}

	if len(trips) > 0 { // Only when new entries fetched from database, cache them.
		// There is no need for API consumer to wait for entries to be cached.
		log.Println("DB Results")
		go func(trips []*pb.TripData, d time.Time) {
			ct.crCache.Cache(trips, pickup)
		}(trips, pickup)
	}

	if err != nil {
		log.Println("error internal handler", err)
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	if len(tt) > 0 {
		trips = append(trips, tt...)
	}
	rs := pb.FetchRs{Pickup: rq.Pickup, Trips: trips}
	return &rs, nil
}

// ClearCache either removes entries for a given pickup date or flushes out all entries in REDIS.
func (ct *CabTrips) ClearCache(ctx context.Context, rq *pb.ClearCacheRq) (*pb.ClearCacheRs, error) {
	// Validate Input Request
	err := ct.Validte(rq)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	if len(rq.Pickup) == 0 {
		err := ct.crCache.Flush()
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, err.Error())
		}
		return &pb.ClearCacheRs{}, nil
	}
	// Validate Step will validate date
	pickup, err := time.Parse("2006-01-02", rq.Pickup)
	err = ct.crCache.Clear(pickup)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &pb.ClearCacheRs{}, nil
}

// Validte verifies input requests for grpc handler service.
func (ct *CabTrips) Validte(rq interface{}) (err error) {
	switch dRq := rq.(type) {
	case *pb.FetchRq:
		err = validation.ValidateStruct(dRq,
			validation.Field(&dRq.Medallions, validation.Required, validation.Length(1, 50)),
			validation.Field(&dRq.Pickup, validation.Required, validation.Date("2006-01-02")),
		)
	case *pb.ClearCacheRq:
		err = validation.ValidateStruct(dRq,
			validation.Field(&dRq.Pickup, validation.Date("2006-01-02")),
		)
	default:
		return status.ErrBadRequest
	}
	return err
}
