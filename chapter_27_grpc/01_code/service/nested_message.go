package service

import (
	"context"
	"log"
	nestedMsg "trip/proto/gen/nested_message"
)

// TripServiceServer is the server API for TripService service.
//type TripServiceServer interface {
//	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
//}

type NestedService struct{}

func (n NestedService) SayHello(ctx context.Context, req *nestedMsg.HelloReq) (*nestedMsg.HelloResp, error) {
	log.Printf("req.Msg: %s, req.Date: %v\n", req.Msg, req.Data)
	return &nestedMsg.HelloResp{
		TraceId: "mememememeda",
	}, nil
}
