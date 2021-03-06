package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/hakobe/grpc-guruguru/go/guruguru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Name           string
	HostPort       string
	PublicHostPort string
	BossHostPort   string
}

func getConfig() *Config {
	memberName := os.Getenv("MEMBER_NAME")
	if memberName == "" {
		memberName = "go"
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
		Name:           memberName,
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
		JoiningMember: &pb.Member{
			Name:     config.Name,
			HostPort: config.PublicHostPort,
		},
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not join: %v", err)
	}
}

func poke(to *Member, config *Config) {
	conn := getConn(to.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewMemberServiceClient(conn)
	res, err := client.Poke(ctx, &pb.PokeRequest{
		FromMember: &pb.Member{
			Name:     config.Name,
			HostPort: config.PublicHostPort,
		},
		Message: "Gopher is the best programming language mascot!!",
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not poke: %v", err)
	}
}

type Server struct {
	Config *Config
	Next   *Member
	lock   sync.RWMutex
}

func (s *Server) SetNext(ctx context.Context, in *pb.SetNextRequest) (*pb.SetNextResponse, error) {
	member := in.GetNextMember()
	if member == nil {
		return &pb.SetNextResponse{Ok: false}, nil
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.Next = &Member{
		Name:     member.Name,
		HostPort: member.HostPort,
	}
	fmt.Printf("Set next to %s(%s)\n", s.Next.Name, s.Next.HostPort)
	return &pb.SetNextResponse{Ok: true}, nil
}

func (s *Server) poke() {
	s.lock.RLock()
	defer s.lock.RUnlock()

	poke(s.Next, s.Config)
}

func (s *Server) Poke(ctx context.Context, in *pb.PokeRequest) (*pb.PokeResponse, error) {
	from := in.GetFromMember()
	fmt.Printf("Got message \"%s\" from %s. Hey %s! \n", in.GetMessage(), from.Name, s.Next.Name)
	go s.poke()
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
	fmt.Printf("Starting... I'm %s\n", config.Name)

	join(config)
	serve(config)
}
