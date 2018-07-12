// Package handler includes implementation of GRPC Service.
// This implementation relies on underlying data-store or cache services implementation pkg/rides/CabRide and pkg/rides/CabRideCache interfaces.
package handler

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

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

	pickup, err := time.Parse("2006-01-02", rq.Pickup)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	tt, notFound, err := ct.crCache.Fetch(rq.Medallions, pickup)
	var trips []*pb.TripData

	switch {
	case len(notFound) == 0:
	case len(notFound) == len(rq.Medallions):
		trips, err = ct.crSrv.Fetch(rq.Medallions, pickup)
	case len(notFound) < len(rq.Medallions):
		trips, err = ct.crSrv.Fetch(notFound, pickup)
	}

	if len(trips) > 0 { // Only when new entries fetched from database, cache them.
		// There is no need for API consumer to wait for entries to be cached.
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
	if len(rq.Pickup) == 0 {
		err := ct.crCache.Flush()
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, err.Error())
		}
		return &pb.ClearCacheRs{}, nil
	}

	pickup, err := time.Parse("2006-01-02", rq.Pickup)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	err = ct.crCache.Clear(pickup)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &pb.ClearCacheRs{}, nil
}
