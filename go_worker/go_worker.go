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

type Member struct {
	Name     string
	HostPort string
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

	c := pb.NewBossServiceClient(conn)

	res, err := c.Join(ctx, &pb.JoinRequest{
		Member: &pb.Member{
			Name:     config.WorkerName,
			HostPort: config.PublicHostPort,
		},
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not join: %v", err)
	}
}

func poke(to *Member) {
	conn := getConn(to.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewMemberServiceClient(conn)
	res, err := client.Poke(ctx, &pb.PokeRequest{
		From: &pb.Member{
			Name:     to.Name,
			HostPort: to.HostPort,
		},
		Message: "Gopher is the best programming language mascot!!",
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not send task: %v", err)
	}

}

type Server struct {
	Config *Config
	Next   *Member
	lock   sync.RWMutex
}

func (s *Server) SetNext(ctx context.Context, in *pb.SetNextRequest) (*pb.SetNextResponse, error) {
	if member := in.GetMember(); member != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.Next = &Member{
			Name:     member.Name,
			HostPort: member.HostPort,
		}
		fmt.Printf("Set next to %s(%s)\n", s.Next.Name, s.Next.HostPort)
		return &pb.SetNextResponse{Ok: true}, nil
	}
	return &pb.SetNextResponse{Ok: false}, nil
}

func (s *Server) poke(ctx context.Context, in *pb.PokeRequest) {
	from := in.GetFrom()
	if from == nil {
		log.Fatalf("could not get from-member")
	}

	fmt.Printf("from:%s -> me:%s -> next:%s\n", from.GetName(), s.Config.WorkerName, s.Next.Name)
	s.lock.RLock()
	defer s.lock.RUnlock()

	poke(s.Next)
}

func (s *Server) Poke(ctx context.Context, in *pb.PokeRequest) (*pb.PokeResponse, error) {
	go s.poke(ctx, in)
	return &pb.PokeResponse{Ok: true}, nil
}

func serve(config *Config) {
	lis, err := net.Listen("tcp", config.HostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMemberServiceServer(s, &Server{
		Config: config,
		Next:   nil,
		lock:   sync.RWMutex{},
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
