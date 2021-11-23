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
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/kaijchen/tracker/track"
	"google.golang.org/grpc"
)

var (
	addr = "localhost:2333"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTrackerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	switch os.Args[1] {
	case "r":
		src, err := strconv.Atoi(os.Args[4])
		if err != nil {
			src = -1
		}
		r, err := c.Report(ctx, &pb.ReportRequest{Key: os.Args[2], Location: os.Args[3], Source: int64(src)})
		if err != nil {
			log.Fatalf("could not report: %v", err)
		}
		log.Printf("Result: %v", r.GetOk())
	case "q":
		r, err := c.Query(ctx, &pb.QueryRequest{Key: os.Args[2]})
		if err != nil {
			log.Fatalf("could not query: %v", err)
		}
		log.Printf("Result: %v, src=%v", r.GetLocation(), r.GetSource())
	}
}
