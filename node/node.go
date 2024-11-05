package main

import (
	"context"
	"sync"

	pb "modname/gRPC"

	//"google.golang.org/grpc"
)


type Node struct {
	pb.ConsensusServer
	pb.ConsensusClient
   
	// Client 
	nodes      map[string]bool
	mu                sync.Mutex
	lamportTime       int64
	nodeId              int
	nextNodeId        int
}

func NewNodeServer() *Node {
	return &Node{
		nodes: make(map[string]bool),
	}
}

//Funktion til SendToken

func (s *Node) Token(ctx context.Context, req *pb.SendToken) (*pb.Void, error) {
	
}