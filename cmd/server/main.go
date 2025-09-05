package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/peer-coding/grpc-comm/internal/http/grpc/server"
	"github.com/peer-coding/grpc-comm/pkg/logger"
)

func main() {
	logger.New()

	go func() {
		addr := fmt.Sprintf("%s:%s", "localhost", "3001")
		server := server.New(addr)

		logger.Log.Info("server is running", "address", addr)

		if err := server.Run(); err != nil {
			logger.Log.Error(err.Error())
			os.Exit(1)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
