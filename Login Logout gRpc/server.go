package main

import (
	"app/authproto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type User struct {
	userName string
	Password string
}

type server struct {
	authproto.UnimplementedSecurityServer
}

var Users = []User{
	{userName: "Amrit", Password: "secert1"},
	{userName: "motti", Password: "qwerty23"},
}

var currUser []User

//implement RPCs

func (s *server) Login(ctx context.Context, loginRequest *authproto.LoginRequest) (loginResponse *authproto.LoginResponse, Error error) {
	for _, user := range Users {
		if user.userName == loginRequest.GetUsername() && user.Password == loginRequest.GetPassword() {
			res := fmt.Sprintf(" The username : %v and Password : %v are correct! Welcome %v", user.userName, user.Password, user.userName)
			currUser = append(currUser, user)
			return &authproto.LoginResponse{Message: res}, nil
		}
		if user.userName == loginRequest.GetUsername() && user.Password != loginRequest.GetPassword() {
			res := fmt.Sprintf(" For username : %v , the password : %v is wrong! Sorry Tru Again!", loginRequest.GetUsername(), loginRequest.GetPassword())
			return &authproto.LoginResponse{Message: res}, nil
		}
	}

	ans := fmt.Sprintf("Sorry the username %v is not found in our database", loginRequest.GetUsername())
	return &authproto.LoginResponse{Message: ans}, nil
}

func (s *server) Logout(ctx context.Context, logout *authproto.LogoutRequest) (logoutResponse *authproto.Logoutresponse, Error error) {

	for i, user := range currUser {
		if user.userName == logout.GetUsername() {

			currUser = append(currUser[:i], currUser[i+1:]...)

			res := fmt.Sprintf("Hey the user : %v is logged out successfully", user.userName)
			return &authproto.Logoutresponse{Message: res}, nil
		}
	}

	res := fmt.Sprintf("Sorry the username : %v is not logged in, Please login first", logout.GetUsername())
	return &authproto.Logoutresponse{Message: res}, nil

}

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcserver := grpc.NewServer()

	authproto.RegisterSecurityServer(grpcserver, &server{})

	fmt.Println("Started the Security Server ....")
	grpcserver.Serve(listener)

}
