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
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/kaijchen/tracker/track"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTrackerServer
	mu      sync.RWMutex
	objects map[string]*rrQueue
	peerID  map[string]int
	peers   []string
}

type rrQueue struct {
	q []int
	i int
}

const (
	defaultConfigPath = "/etc/tracker/config.json"
)

type config struct {
	Port     int      `json:"port"`
	Registry []string `json:"registry"`
}

var (
	registryIDs []int
)

func getConfig(path string) (config, error) {
	cfg := config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func (rq *rrQueue) get() int {
	if len(rq.q) == 0 {
		return -1
	}
	rq.i = (rq.i + 1) % len(rq.q)
	return rq.q[rq.i]
}

func (rq *rrQueue) put(x int) {
	rq.q = append(rq.q, x)
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	log.Printf("Fetching: %v", in.GetKey())
	var loc string
	s.mu.RLock()
	rrq := s.objects[in.GetKey()]
	if rrq == nil {
		rrq = &rrQueue{}
		rrq.q = append(rrq.q, registryIDs...)
		s.objects[in.GetKey()] = rrq
	}
	id := rrq.get()
	if id >= 0 {
		loc = s.peers[id]
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
	rrq := s.objects[in.GetKey()]
	if rrq == nil {
		rrq = &rrQueue{}
		rrq.q = append(rrq.q, registryIDs...)
		s.objects[in.GetKey()] = rrq
	}
	rrq.put(s.peerID[in.GetLocation()])
	s.mu.Unlock()
	return &pb.ReportReply{Ok: true}, nil
}

func main() {
	cfg, err := getConfig(defaultConfigPath)
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := server{objects: make(map[string]*rrQueue), peerID: make(map[string]int)}
	for i, r := range cfg.Registry {
		srv.peers = append(srv.peers, r)
		srv.peerID[r] = i
		registryIDs = append(registryIDs, i)
	}
	pb.RegisterTrackerServer(s, &srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
