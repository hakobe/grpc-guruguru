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

type Worker struct {
	Name     string
	HostPort string
}

type Server struct {
	Workers map[string]*Worker
	lock    sync.RWMutex
}

func newServer() *Server {
	return &Server{
		Workers: make(map[string]*Worker),
		lock:    sync.RWMutex{},
	}
}

func (s *Server) Join(ctx context.Context, in *pb.JoinRequest) (*pb.Res, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Workers[in.GetName()] = &Worker{
		Name:     in.GetName(),
		HostPort: in.GetHostPort(),
	}

	fmt.Printf("%s is joined.\n", in.GetName())

	return &pb.Res{Ok: true}, nil
}

func (s *Server) GetWorkers() []*Worker {
	s.lock.RLock()
	defer s.lock.RUnlock()

	workers := []*Worker{}
	for _, worker := range s.Workers {
		workers = append(workers, worker)
	}
	shuffle(workers)

	return workers
}

func shuffle(data []*Worker) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
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

func serve(server *Server, hostPort string) {
	lis, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBossServer(grpcServer, server)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getConn(hostPort string) *grpc.ClientConn {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func sendTask(worker *Worker, fromName string) {
	conn := getConn(worker.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewWorkerClient(conn)
	res, err := client.AcceptTask(ctx, &pb.Task{
		FromName: fromName,
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not send task: %v", err)
	}
}

func setNext(worker *Worker, nextWorker *Worker) {
	conn := getConn(worker.HostPort)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewWorkerClient(conn)
	res, err := client.SetNext(ctx, &pb.Next{
		NextName:     nextWorker.Name,
		NextHostPort: nextWorker.HostPort,
	})
	if err != nil || !res.GetOk() {
		log.Fatalf("could not set next: %v", err)
	}
}

func startTask(server *Server) {
	time.Sleep(5 * time.Second)
	fmt.Println("Workers are gathered.")

	workers := server.GetWorkers()
	n := len(workers)
	for i := 0; i < n; i++ {
		worker := workers[i]
		nextWorker := workers[(i+1)%n]

		fmt.Printf("Setting worker %s(%s) -> %s(%s)\n", worker.Name, worker.HostPort, nextWorker.Name, nextWorker.HostPort)
		setNext(worker, nextWorker)
	}

	if len(workers) > 0 {
		sendTask(workers[0], "boss")
	}
}

func main() {
	config := getConfig()
	fmt.Println("Starting... I'm BOSS.")

	server := newServer()

	go startTask(server)
	serve(server, config.HostPort)
}
