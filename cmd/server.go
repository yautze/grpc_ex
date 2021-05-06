package cmd

import (
	"fmt"
	"log"
	"net"

	"grpc_ex/controller"
	pb "grpc_ex/protobuf"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// port -
var port = ""

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "p", "3333", "server")
}

func start() {
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	// 啟動一個監聽的server
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("start listener failed: %v", err)
		return
	}

	log.Println("server listening on", addr)

	// 啟動 grpc server
	grpcSrv := grpc.NewServer()
	// 將實作好的server註冊到grpc service
	server := controller.NewServer()
	pb.RegisterHelloServiceServer(grpcSrv, server)

	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("start server failed: %v", err)
		return
	}
}
