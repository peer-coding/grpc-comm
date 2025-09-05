package server

import (
	"net"
	"os"

	"github.com/peer-coding/grpc-comm/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	lis net.Listener
	srv *grpc.Server
}

func New(addr string) Server {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	return Server{
		lis: lis,
		srv: server,
	}
}

func (s *Server) Run() error {
	if err := s.srv.Serve(s.lis); err != nil {
		return err
	}

	return nil
}
