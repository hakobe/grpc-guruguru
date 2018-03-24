package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/hakobe/grpc-guruguru/boss/guruguru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Worker struct {
	Name     string
	HostPort string
}

type Server struct {
	Workers map[string]*Worker
	lock    sync.RWMutex
}

func (s *Server) Join(ctx context.Context, in *pb.JoinRequest) (*pb.Res, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Workers[in.GetName()] = &Worker{
		Name:     in.GetName(),
		HostPort: in.GetHostPort(),
	}
	for k, v := range s.Workers {
		fmt.Printf("%v = %v\n", k, v.HostPort)
	}
	return &pb.Res{Ok: true}, nil
}

type Config struct {
	HostPort string
}

func getConfig() *Config {
	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		hostPort = "0.0.0.0:5000"
	}

	return &Config{
		HostPort: hostPort,
	}
}

func main() {
	fmt.Println("Starting... I'm BOSS.")

	config := getConfig()

	lis, err := net.Listen("tcp", config.HostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := &Server{
		Workers: make(map[string]*Worker),
		lock:    sync.RWMutex{},
	}
	pb.RegisterBossServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
