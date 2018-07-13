package handler_test

import (
	"encoding/json"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/govinda-attal/cabride-api/pkg/rides/pb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/govinda-attal/cabride-api/handler"

	"github.com/govinda-attal/cabride-api/test/mocks"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("CabTrips", func() {

	Describe("Basic Behaviour", func() {
		var (
			timestamp = func(tStr string) *timestamp.Timestamp {
				t, _ := time.Parse("2006-01-02T15:04:05Z", tStr)
				ts, _ := ptypes.TimestampProto(t)
				return ts
			}

			pickup = "2013-12-01"

			medallions = []string{"D7D598CD99978BD012A87A76A7C891B7"}

			trips = []*pb.TripData{
				&pb.TripData{
					Medallion:      "D7D598CD99978BD012A87A76A7C891B7",
					HackLicense:    "82F90D5EFE52FDFD2FDEC3EAD6D5771D",
					VendorID:       "VTS",
					RateCode:       1,
					Pickup:         timestamp("2013-12-01T00:13:00Z"),
					Dropoff:        timestamp("2013-12-01T00:31:00Z"),
					PassengerCount: 1,
					TripTime:       1080,
					TripDistance:   3.79,
				},
			}
		)

		BeforeEach(func() {

		})

		Context("When given a valid travel date and medallion", func() {
			mockCRStore := &mocks.CabrideMockDS{}
			mockCRCache := &mocks.CabrideMockCache{}
			mockCRStore.FetchCall.Returns.Trips = trips
			mockCRCache.FetchCall.Returns.NotFoundMedallions = medallions
			It("must return cab trips for given pickup date and no error must have occurred", func() {
				fetchGRPCRq := &pb.FetchRq{Pickup: pickup, Medallions: medallions}
				expectedGRPCRs := &pb.FetchRs{Pickup: pickup, Trips: trips}
				handler := NewCabTripsHandler(mockCRStore, mockCRCache)
				grpcRs, err := handler.Fetch(context.Background(), fetchGRPCRq)
				Expect(err).NotTo(HaveOccurred())
				rsBytes, _ := json.Marshal(grpcRs)
				expRsBytes, _ := json.Marshal(expectedGRPCRs)
				Expect(rsBytes).To(MatchJSON(expRsBytes))
			})
		})

		Context("When given a valid travel date, medallion and explict flag to not use cache", func() {
			mockCRStore := &mocks.CabrideMockDS{}
			mockCRCache := &mocks.CabrideMockCache{}
			mockCRStore.FetchCall.Returns.Trips = trips
			mockCRCache.FetchCall.Returns.NotFoundMedallions = medallions
			It("must not call/fetch from cache, return cab trips for given pickup date and no error must have occurred", func() {
				fetchGRPCRq := &pb.FetchRq{Pickup: pickup, Medallions: medallions, NoCache: true}
				expectedGRPCRs := &pb.FetchRs{Pickup: pickup, Trips: trips}
				handler := NewCabTripsHandler(mockCRStore, mockCRCache)
				grpcRs, err := handler.Fetch(context.Background(), fetchGRPCRq)
				Expect(err).NotTo(HaveOccurred())
				rsBytes, _ := json.Marshal(grpcRs)
				expRsBytes, _ := json.Marshal(expectedGRPCRs)
				Expect(rsBytes).To(MatchJSON(expRsBytes))
				Expect(mockCRCache.FetchCall.Receives).To(BeZero())
			})
		})
	})
})
