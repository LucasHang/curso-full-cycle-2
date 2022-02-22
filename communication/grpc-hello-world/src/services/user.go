package services

import (
	"context"
	"fmt"
	"time"
	"log"
	"io"
	"github.com/LucasHang/curso-full-cycle-2/communication/grpc-hello-world/src/pb"
);

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{};
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Insert - Database
	fmt.Println(req.GetName());

	return &pb.User{
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil;
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User: &pb.User{},
	});

	time.Sleep(time.Second * 3);

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	});

	time.Sleep(time.Second * 3);

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	});

	time.Sleep(time.Second * 3);

	return nil;
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{};

	for {
		req, err := stream.Recv();

		if(err != nil){
			if(err == io.EOF){
				return stream.SendAndClose(&pb.Users{
					User: users,
				})
			}
			
			log.Fatalf("Error receiving stream: %v", err);
		}

		users = append(users, &pb.User{
			Id: req.GetId(),
			Name: req.GetName(),
			Email: req.GetEmail(),
		});

		fmt.Println("Adding: ", req.GetName());
	}
}