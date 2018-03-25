package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/hakobe/grpc-guruguru/boss/guruguru"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

func poke(from *Member, to *Member) {
	conn := getConn(to.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewMemberServiceClient(conn)
	res, err := client.Poke(ctx, &pb.PokeRequest{
		From: &pb.Member{
			Name:     from.Name,
			HostPort: from.HostPort,
		},
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not send poke: %v", err)
	}
}

func setNext(to *Member, next *Member) {
	conn := getConn(to.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewMemberServiceClient(conn)
	res, err := client.SetNext(ctx, &pb.SetNextRequest{
		Member: &pb.Member{
			Name:     next.Name,
			HostPort: next.HostPort,
		},
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not set next: %v", err)
	}
}

type Server struct {
	Members map[string]*Member
	lock    sync.RWMutex
}

func newServer() *Server {
	return &Server{
		Members: make(map[string]*Member),
		lock:    sync.RWMutex{},
	}
}

func (s *Server) Join(ctx context.Context, in *pb.JoinRequest) (*pb.JoinResponse, error) {
	member := in.GetMember()
	if member == nil {
		log.Fatalf("could not get joining member")
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.Members[member.GetName()] = &Member{
		Name:     member.GetName(),
		HostPort: member.GetHostPort(),
	}

	fmt.Printf("%s is joined.\n", member.GetName())

	return &pb.JoinResponse{Ok: true}, nil
}

func (s *Server) GetMembers() []*Member {
	s.lock.RLock()
	defer s.lock.RUnlock()

	members := []*Member{}
	for _, member := range s.Members {
		members = append(members, member)
	}
	shuffle(members)

	return members
}

func shuffle(data []*Member) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func serve(server *Server, hostPort string) {
	lis, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBossServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startTask(server *Server) {
	time.Sleep(5 * time.Second)
	fmt.Println("Workers are gathered.")

	members := server.GetMembers()
	n := len(members)
	for i := 0; i < n; i++ {
		member := members[i]
		nextMember := members[(i+1)%n]

		fmt.Printf("Setting worker %s(%s) -> %s(%s)\n", member.Name, member.HostPort, nextMember.Name, nextMember.HostPort)
		setNext(member, nextMember)
	}

	if len(members) > 0 {
		poke(&Member{Name: "boss", HostPort: ""}, members[0])
	}
}

func main() {
	config := getConfig()
	fmt.Println("Starting... I'm BOSS.")

	server := newServer()

	go startTask(server)
	serve(server, config.HostPort)
}
