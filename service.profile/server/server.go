package server

import (
	pb "github.com/opiskull/red-panda/service.profile/proto"
	"golang.org/x/net/context"
)

// NewServer creates a new authServer and loads the customers from
// JSON file into customers map
func NewServer(conf string) *ProfileServer {
	s := &ProfileServer{}
	return s
}

// authServer struct w/ customers map
type ProfileServer struct {
	serverName string
}

// VerifyToken finds a customer by authentication token.
func (s *ProfileServer) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {

	reply := &pb.ProfileReply{
		Id:    "hamster",
		Email: "hamster",
	}

	return reply, nil
}
