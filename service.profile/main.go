package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	// "time"

	pb "github.com/opiskull/red-panda/service.profile/proto"
	"github.com/opiskull/red-panda/service.profile/server"
	// trace "github.com/harlow/go-micro-services/api.trace/client"

	"google.golang.org/grpc"
)

func main() {
	var (
		port = flag.Int("port", 10001, "The server port")
		conf = flag.String("config", "etcd://localhost:1111", "A json file containing a list of customers")
	)

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProfileServer(grpcServer, server.NewServer(*conf))
	grpcServer.Serve(lis)
	defer grpcServer.Stop()
}
