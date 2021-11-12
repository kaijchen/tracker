/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/kaijchen/tracker/track"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("config", 2333, "The tracker port")
)

type server struct {
	pb.UnimplementedTrackerServer
	mu      sync.RWMutex
	objects map[string][]int
	peerID  map[string]int
	peers   []string
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	log.Printf("Fetching: %v", in.GetKey())
	var loc string
	s.mu.RLock()
	peerIDs := s.objects[in.GetKey()]
	if len(peerIDs) > 0 {
		loc = s.peers[peerIDs[0]]
	}
	s.mu.RUnlock()
	return &pb.QueryReply{Location: loc}, nil
}

func (s *server) Report(ctx context.Context, in *pb.ReportRequest) (*pb.ReportReply, error) {
	log.Printf("Report: %v at %v", in.GetKey(), in.GetLocation())
	s.mu.Lock()
	if _, ok := s.peerID[in.GetLocation()]; !ok {
		s.peerID[in.GetLocation()] = len(s.peers)
		s.peers = append(s.peers, in.GetLocation())
	}
	s.objects[in.GetKey()] = append(s.objects[in.GetKey()], s.peerID[in.GetLocation()])
	s.mu.Unlock()
	return &pb.ReportReply{Ok: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTrackerServer(s, &server{objects: make(map[string][]int), peerID: make(map[string]int)})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}