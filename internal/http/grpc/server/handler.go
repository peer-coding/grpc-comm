package server

import (
	"context"
	"os"
	"time"

	"github.com/peer-coding/grpc-comm/api/proto/pb"
	"github.com/peer-coding/grpc-comm/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedHandshakeServiceServer
}

func (h *Handler) OneWayHandshake(ctx context.Context, req *pb.HandshakeRequest) (*pb.HandshakeResponse, error) {
	logger.Log.Info("request received", "sender", req.Sender, "message", req.Message)

	hostname, err := os.Hostname()

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "unable to get host")
	}

	return &pb.HandshakeResponse{
		Receiver:   hostname,
		Status:     pb.CommunicationType_Success,
		ReceivedAt: timestamppb.New(time.Now()),
	}, nil
}
