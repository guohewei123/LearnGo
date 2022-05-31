package service

import (
	"context"
	"log"
	timestamp "trip/proto/gen/timestamp"
)

type TimestampService struct{}

func (t TimestampService) AddTime(ctx context.Context, req *timestamp.TimestampReq) (*timestamp.TimestampResp, error) {
	log.Printf("map data: %v\n", req)
	return &timestamp.TimestampResp{TraceId: "map123"}, nil
}

