package main

import (
	"context"
	"time"

	pb "github.com/brotherlogic/pingspeed/proto"
)

func (s *Server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{ReceivedTime: time.Now().UnixNano(), ReturnedTime: time.Now().UnixNano()}, nil
}
