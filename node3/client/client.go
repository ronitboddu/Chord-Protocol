package client

import (
	"Test1/pb"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FindKey(succ_ip_addr string, succ_port string, key string) (*pb.ResponseNode, error) {
	conn, err := grpc.Dial(succ_ip_addr+":"+succ_port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Lookup(ctx, &pb.Key{Key: key})

	if err != nil {
		panic(err)
	}
	return res, nil

}
