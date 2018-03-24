package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/hakobe/grpc-guruguru/go_worker/guruguru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func (s *Server) SetNext(ctx context.Context, in *pb.Next) (*pb.Res, error) {
	return &pb.Res{Ok: true}, nil
}
func (s *Server) AcceptTask(ctx context.Context, in *pb.Task) (*pb.Res, error) {
	return &pb.Res{Ok: true}, nil
}

type Config struct {
	WorkerName     string
	HostPort       string
	PublicHostPort string
	BossHostPort   string
}

func getConfig() *Config {
	workerName := os.Getenv("WORKER_NAME")
	if workerName == "" {
		workerName = "go_worker"
	}
	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		hostPort = "0.0.0.0:5000"
	}
	publicHostPort := os.Getenv("PUBLIC_HOST_PORT")
	if publicHostPort == "" {
		publicHostPort = "localhost:5000"
	}
	bossHostPort := os.Getenv("BOSS_HOST_PORT")
	if bossHostPort == "" {
		bossHostPort = "localhost:5001"
	}

	return &Config{
		WorkerName:     workerName,
		HostPort:       hostPort,
		PublicHostPort: publicHostPort,
		BossHostPort:   bossHostPort,
	}
}

func main() {
	config := getConfig()

	fmt.Printf("Starting... I'm %s\n", config.WorkerName)

	conn, err := grpc.Dial(config.BossHostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBossClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Join(ctx, &pb.JoinRequest{
		Name:     "go",
		HostPort: config.PublicHostPort,
	})
	if err != nil {
		log.Fatalf("could not join: %v", err)
	}

	lis, err := net.Listen("tcp", config.HostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWorkerServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
