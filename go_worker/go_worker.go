package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/hakobe/grpc-guruguru/go_worker/guruguru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Config       *Config
	NextName     string
	NextHostPort string
	lock         sync.RWMutex
}

func (s *Server) SetNext(ctx context.Context, in *pb.Next) (*pb.Res, error) {
	fmt.Printf("Set next to %s(%s)\n", in.NextName, in.NextHostPort)
	s.lock.Lock()
	defer s.lock.Unlock()

	s.NextName = in.GetNextName()
	s.NextHostPort = in.GetNextHostPort()

	return &pb.Res{Ok: true}, nil
}

func (s *Server) acceptTask(ctx context.Context, in *pb.Task) {
	fmt.Printf("from:%s -> me:%s -> next:%s\n", in.FromName, s.Config.WorkerName, s.NextName)
	s.lock.RLock()
	defer s.lock.RUnlock()

	conn := getConn(s.NextHostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewWorkerClient(conn)
	res, err := client.AcceptTask(ctx, &pb.Task{
		FromName: s.Config.WorkerName,
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not send task: %v", err)
	}
}

func (s *Server) AcceptTask(ctx context.Context, in *pb.Task) (*pb.Res, error) {
	go s.acceptTask(ctx, in)
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

func getConn(hostPort string) *grpc.ClientConn {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func join(config *Config) {
	conn := getConn(config.BossHostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewBossClient(conn)

	res, err := c.Join(ctx, &pb.JoinRequest{
		Name:     config.WorkerName,
		HostPort: config.PublicHostPort,
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not join: %v", err)
	}
}

func serve(config *Config) {
	lis, err := net.Listen("tcp", config.HostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWorkerServer(s, &Server{
		Config:       config,
		NextName:     "",
		NextHostPort: "",
		lock:         sync.RWMutex{},
	})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	config := getConfig()
	fmt.Printf("Starting... I'm %s\n", config.WorkerName)

	join(config)
	serve(config)
}
