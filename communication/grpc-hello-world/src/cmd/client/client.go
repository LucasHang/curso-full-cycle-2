package main;

import (
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/LucasHang/curso-full-cycle-2/communication/grpc-hello-world/src/pb"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure());
	if(err != nil){
		log.Fatalf("Could not connect to gRPC server: %v", err);
	}

	defer connection.Close();

	client := pb.NewUserServiceClient(connection);
	AddUser(client);
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Lucas",
		Email: "lucas@gmail.com",
	};

	res, err := client.AddUser(context.Background(), req);

	if(err != nil){
		log.Fatalf("Could not make gRPC request: %v", err);
	}

	log.Println(res);
}