package main;

import (
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/LucasHang/curso-full-cycle-2/communication/grpc-hello-world/src/pb"
	"github.com/LucasHang/curso-full-cycle-2/communication/grpc-hello-world/src/services"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051");
	if(err != nil){
		log.Fatalf("Could not connect: %v", err);
	}

	grpcServer := grpc.NewServer();
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService());
	reflection.Register(grpcServer);

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err);
	}
}