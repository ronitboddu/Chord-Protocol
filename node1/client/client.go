package client

import (
	"Test1/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FindKey(succ_ip_addr string, succ_port string, key string) (*pb.ResponseNode, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, succ_ip_addr+":"+succ_port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := c.RPCLookup(ctx, &pb.Key{Key: key})
	if err != nil {
		panic(err)
	}
	fmt.Println("here")
	return res, nil

}
