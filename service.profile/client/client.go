package client

import (
	pb "github.com/opiskull/red-panda/service.profile/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ProfileClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewProfileClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c Client) GetProfile(ctx context.Context) (*pb.ProfileReply, error) {
	request := &pb.ProfileRequest{}
	return c.client.GetProfile(ctx, request)
}

func (c Client) Close() error {
	return c.conn.Close()
}
