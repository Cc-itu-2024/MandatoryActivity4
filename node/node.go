package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb "modname/gRPC"

	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedConsensusServer
	client pb.ConsensusClient
	port int
	ports []int
	mu sync.Mutex
	nodeId int
	nextNodeId int
	hasToken bool
}

func NewNodeServer(nodeId, nextNodeId, port int, ports []int, hasToken bool) *Node {
	return &Node{
		nodeId: nodeId,
		nextNodeId: nextNodeId,
		port: port,
		ports: ports,
		hasToken: hasToken,
	}
}

func (s *Node) Token(ctx context.Context, req *pb.SendToken) (*pb.Void, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Node %d received the token from Node %d", s.nodeId, req.NodeId)
	s.hasToken = true
	s.enterCriticalSection()

	return &pb.Void{}, nil
}

func (s *Node) enterCriticalSection() {
	if s.hasToken {
		log.Printf("Node %d entered critical section\n", s.nodeId)
		time.Sleep(2 * time.Second)

		log.Printf("Node %d exited critical section\n", s.nodeId)
		s.passToken()
	}
}

func (s *Node) passToken() {

	if !s.hasToken {
		log.Printf("Node %d does not have the token to pass.\n", s.nodeId)
		return
	}
	s.hasToken = false

	log.Printf("Node %d passed the token to Node %d\n", s.nodeId, s.nextNodeId)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", s.ports[s.nextNodeId]), grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to next node on port %d: %v", s.ports[s.nextNodeId], err)
		return
	}
	defer conn.Close()

	client := pb.NewConsensusClient(conn)
	_, err = client.Token(context.TODO(), &pb.SendToken{NodeId: int32(s.nodeId)})
	if err != nil {
		log.Printf("Failed to pass token from Node %d to Node %d: %v", s.nodeId, s.nextNodeId, err)
	} else {
		log.Printf("Node %d passed the token to Node %d\n", s.nodeId, s.nextNodeId)
	}
}

func (s *Node) StartServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", s.port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterConsensusServer(grpcServer, s)

	log.Printf("Node %d started server on port %d", s.nodeId, s.port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", s.port, err)
	}
}

func main() {
	logFile, err := os.OpenFile("consensus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	numNodes := 3
	basePort := 5100
	nodes := make([]*Node, numNodes)
	ports := make([]int, numNodes)

	for i := 0; i < numNodes; i++ {
		port := basePort + i*2
		ports[i] = port
		log.Printf("Assigned port %d to Node %d", port, i)
	}

	for i := 0; i < numNodes; i++ {
		hasToken := (i == 0)
		nextNodeId := (i + 1) % numNodes

		nodes[i] = NewNodeServer(i, nextNodeId, ports[i], ports, hasToken)
		go nodes[i].StartServer()
	}

	time.Sleep(1 * time.Second)

	if nodes[0].hasToken {
		log.Printf("Started token passing from Node %d\n", nodes[0].nodeId)
		nodes[0].passToken()
	} else {
		log.Printf("Node %d does not have the token to start\n", nodes[0].nodeId)
	}

	select {}
}
