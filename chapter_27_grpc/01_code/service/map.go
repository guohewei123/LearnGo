package service

import (
	"context"
	"log"
	mapbp "trip/proto/gen/map"
)

type MapService struct{}

func (m MapService) MapUsage(ctx context.Context, req *mapbp.MapReq) (*mapbp.MapResp, error) {
	log.Printf("map data: %v\n", req.MapData)
	return &mapbp.MapResp{TraceId: "map123"}, nil
}
