package client

import (
	"context"
	"net"
	"os"

	"github.com/peer-coding/grpc-comm/api/proto/pb"
	"github.com/peer-coding/grpc-comm/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	lis net.Listener
	cl  *grpc.ClientConn
}

func New(addr, targetAddr string) (Client, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	client, err := grpc.NewClient(targetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Client{}, err
	}

	return Client{
		lis: lis,
		cl:  client,
	}, nil
}

func (c *Client) Handshake(in *pb.HandshakeRequest) error {
	hsCl := pb.NewHandshakeServiceClient(c.cl)

	res, err := hsCl.OneWayHandshake(context.Background(), in)

	if err != nil {
		return err
	}

	logger.Log.Info("response",
		"receiver", res.Receiver,
		"status", res.Status,
		"messages", res.Messages,
		"received_at", res.ReceivedAt,
	)

	return nil
}
