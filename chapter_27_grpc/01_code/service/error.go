package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	errortestpb "trip/proto/gen/errortest"
)

type ErrorService struct{}

func (e ErrorService) ErrorUsage(ctx context.Context, empty *empty.Empty) (*errortestpb.RespData, error) {
	//return nil, status.Error(codes.Unauthenticated, "鉴权失败")
	time.Sleep(time.Second * 5)
	return nil, status.New(codes.Unauthenticated, "鉴权失败").Err()
}
