package cmd

import (
	"context"
	"fmt"
	"github.com/3vilM33pl3/hexclient/internal/pkg/hexcloud"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	grpcClient hexcloud.HexagonServiceClient
}

func NewClient(serverAddr string) (c *Client, err error) {
	c = &Client{}
	err = c.Connect(serverAddr)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	return
}

func (c *Client) Status() (message string, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	status, err := c.grpcClient.GetStatus(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return status.Msg, nil
}

func (c *Client) Connect(serverAddr string) (err error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Unable to connect: %s", err)
	}
	c.grpcClient = hexcloud.NewHexagonServiceClient(conn)

	return nil
}
