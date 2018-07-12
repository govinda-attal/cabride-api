package rideds

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"

	"github.com/govinda-attal/cabride-api/internal/provider"
	"github.com/govinda-attal/cabride-api/pkg/core/status"
	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// CabRideCache is an implementation of pkg/ride/CabRide.
type CabRideCache struct {
	redCon redis.Conn
}

// NewCabRideCache returns new instance of cache store service.
func NewCabRideCache() *CabRideCache {
	return &CabRideCache{provider.Cache()}
}

// Fetch return trip data information for given set of medallions on a given pickup date.
// It also returns medallions for which entries were not found in cache if any.
// REDIS CMD SUNION keys... is used to get all records in cache SET for a given medallion on a given date
// For REDIS set name <DATE>_<MEDALLION> convention is followed
func (crc *CabRideCache) Fetch(medallions []string, d time.Time) ([]*pb.TripData, []string, error) {

	// Keys are generated as slice of strings in format <DATE>_<MEDALLION>
	keys := compKeys(d, medallions)

	values, err := redis.ByteSlices(crc.redCon.Do("SUNION", keys...))
	if err != nil {
		return nil, nil, status.ErrInternal.WithMessage(err.Error())
	}

	// Start with assumption that none were found in the cache.
	notFound := make([]string, len(medallions))
	copy(notFound, medallions)
	sort.Strings(notFound)

	trips := []*pb.TripData{}

	// This loop is to unmarshal the []byte tripData
	for _, v := range values {
		trip := pb.TripData{}
		err := proto.Unmarshal(v, &trip)
		if err != nil {
			return nil, nil, status.ErrInternal.WithMessage(err.Error())
		}
		// As for a given medallion an entry in the cache was fetched, remove the medallion from unfound list
		if i := sort.SearchStrings(notFound, trip.Medallion); i < len(notFound) {
			notFound = append(notFound[:i], notFound[i+1:]...)
		}
		trips = append(trips, &trip)
	}
	return trips, notFound, nil
}

// Cache saves trip data information for given set of medallions on a given pickup date.
// Redis SADD Command will be used to add entries to given set/key for each medallion.
// For REDIS set name <DATE>_<MEDALLION> convention is followed
func (crc *CabRideCache) Cache(trips []*pb.TripData, d time.Time) error {

	dPrefix := time.Time(d).Format("2006-01-02")

	chnTD := make(chan *pb.TripData)

	go func() {
		defer close(chnTD)
		for _, t := range trips {
			chnTD <- t
		}
	}()

	go func() {
		for t := range chnTD {
			if kvpair, err := compKeyValPair(dPrefix, t); err == nil {
				crc.redCon.Do("SADD", kvpair...)
			}
		}
	}()
	return nil
}

// compKeys return keys that can be used to MGET values from cache : []{<DATE>_<MEDALLION>}.
func compKeys(d time.Time, medallions []string) []interface{} {
	prefix := time.Time(d).Format("2006-01-02")
	var args []interface{}
	for _, k := range medallions {
		args = append(args, strings.Join([]string{prefix, k}, "_"))
	}
	return args
}

// compKeyValPair return composite set-name/key & value pair for given date and trip data.
// These values will be SADD in the cache : {<DATE>_<MEDALLION>, []byte, ...}.
func compKeyValPair(dPrefix string, t *pb.TripData) ([]interface{}, error) {
	kvpair := make([]interface{}, 2)

	kvpair[0] = strings.Join([]string{dPrefix, t.Medallion}, "_")
	b, err := proto.Marshal(t)
	if err != nil {
		return nil, status.ErrInternal.WithMessage(err.Error())
	}
	kvpair[1] = b
	return kvpair, nil
}

// Clear removes entries for a given date
func (crc *CabRideCache) Clear(d time.Time) error {
	var pattern interface{}
	pattern = fmt.Sprintf("%s*", time.Time(d).Format("2006-01-02"))
	keys, err := redis.Values(crc.redCon.Do("KEYS", []interface{}{pattern}...))
	if len(keys) == 0 {
		return nil
	}
	_, err = crc.redCon.Do("DEL", keys...)
	if err != nil {
		return status.ErrInternal.WithMessage(err.Error())
	}
	return nil
}

// Flush resets the cache
func (crc *CabRideCache) Flush() error {
	_, err := crc.redCon.Do("FLUSHDB")
	if err != nil {
		return status.ErrInternal.WithMessage(err.Error())
	}
	return nil
}
