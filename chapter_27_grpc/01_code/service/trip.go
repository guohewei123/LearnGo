package service

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	trippb "trip/proto/gen/trip"
)
import "context"

// TripServiceServer is the server API for TripService service.
//type TripServiceServer interface {
//	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
//}

type TripService struct{}

func (s *TripService) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println(md)

		if key, ok := md["key1"]; ok {
			fmt.Println(key)
		}
		fmt.Println()
		for k, v := range md {
			fmt.Printf("key: %s, val: %v\n", k, v)
		}
	}

	return &trippb.GetTripResponse{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "shanghai",
			End:         "beijing",
			DurationSec: 1500,
			FeeCent:     1200,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  31,
				Longitude: 121,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 121.1,
				},
				{
					Latitude:  31,
					Longitude: 121.1,
				},
			},
			Status: trippb.TripStatus_PAID,
		},
	}, nil
}
