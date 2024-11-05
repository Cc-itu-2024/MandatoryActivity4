package main

import (
	"context"
	"log"
	"os"
	"sync"

	pb "modname/gRPC"
	//"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedConsensusServer                    //server
	client                          pb.ConsensusClient // client
	Ports                           []int
	mu                              sync.Mutex
	lamportTime                     int64
	nodeId                          int
	nextNodeId                      int
}

func NewNodeServer(nodeId int, ports []int) *Node {
	return &Node{
		nodeId: nodeId,
		Ports:  ports,
	}
}

// Funktion til SendToken
func (s *Node) Token(ctx context.Context, req *pb.SendToken) (*pb.Void, error) {

	return &pb.Void{}, nil
}

func main() {
	logFile, err := os.OpenFile("consensus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	numNodes := 3
	basePort := 5000

	nodes := make([]*Node, numNodes)

	for i := 0; i < numNodes; i++ {
		ports := []int{basePort + i, basePort + i + 1}

		nodes[i] = NewNodeServer(i, ports)
	}

	if len(nodes[numNodes-1].Ports) > 0 && len(nodes[0].Ports) > 0 {
		nodes[numNodes-1].Ports = []int{nodes[numNodes-1].Ports[0], nodes[0].Ports[0]}
	}

	for _, node := range nodes {
		log.Printf("Node %d started with ports: %v", node.nodeId, node.Ports)
	}
}
