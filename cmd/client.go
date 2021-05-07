package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc_ex/protobuf"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	connPort = ""
	mode     = ""
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		client()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&connPort, "connPort", "c", "3333", "conn port")
	clientCmd.Flags().StringVarP(&mode, "mode", "m", "1", "test mode")
}

func client() {
	addr := fmt.Sprintf("127.0.0.1:%s", connPort)

	// connect to grpc server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	// grpc client
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch mode {
	case "1":
		unary(ctx, c)
	}
}

func unary(ctx context.Context, c pb.HelloServiceClient) {
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "YauTz"})
	if err != nil {
		log.Fatalf("could not get nonce: %v", err)
		return
	}

	fmt.Println("Response:", r.GetReply())
}
