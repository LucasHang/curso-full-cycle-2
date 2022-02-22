package main;

import (
	"log"
	"time"
	"io"
	"fmt"
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
	// AddUser(client);
	// AddUserVerbose(client);
	AddUsers(client);
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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Lucas",
		Email: "lucas@gmail.com",
	};

	responseStream, err := client.AddUserVerbose(context.Background(), req);

	if(err != nil){
		log.Fatalf("Could not make gRPC request: %v", err);
	}

	for {
		stream, err := responseStream.Recv();

		if(err != nil){
			if(err == io.EOF){
				break;
			}

			log.Fatalf("Could not receive the message: %v", err);
		}

		log.Println("Status: ", stream.Status);
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "l1",
			Name: "lucas 1",
			Email: "lucas1@gmail.com",
		},
		&pb.User{
			Id: "l2",
			Name: "lucas 2",
			Email: "lucas2@gmail.com",
		},
		&pb.User{
			Id: "l3",
			Name: "lucas 3",
			Email: "lucas3@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background());
	if(err != nil){
		log.Fatalf("Error creating request: %v", err);
	}

	for _, req := range reqs {
		stream.Send(req);
		time.Sleep(time.Second * 3);
	}

	res, err := stream.CloseAndRecv();
	if(err != nil){
		log.Fatalf("Error receiving response: %v", err);
	}

	fmt.Println(res);
}