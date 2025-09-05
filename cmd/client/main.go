package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/peer-coding/grpc-comm/api/proto/pb"
	"github.com/peer-coding/grpc-comm/internal/http/grpc/client"
	"github.com/peer-coding/grpc-comm/pkg/logger"
)

func main() {
	logger.New()

	go func() {
		addr := fmt.Sprintf("%s:%s", "localhost", "3002")
		serverAddr := fmt.Sprintf("%s:%s", "localhost", "3001")

		cl, err := client.New(addr, serverAddr)
		if err != nil {
			logger.Log.Error("error initializing client", "error", err)
			os.Exit(1)
		}

		logger.Log.Info("server is running", "address", addr)

		handshakes := []*pb.HandshakeRequest{
			{
				Sender: "Arthur",
			},
			{
				Sender: "Alan",
			},
			{
				Sender: "Gustavo",
			},
		}

		for _, hs := range handshakes {
			logger.Log.Info("handshaking", "sender", hs.Sender)

			if err := cl.Handshake(hs); err != nil {
				logger.Log.Error("handshake error", "error", err)
			}

			logger.Log.Info("handshaked successfully", "sender", hs.Sender)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
