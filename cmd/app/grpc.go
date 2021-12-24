package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "payment.system.com/proto"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	app *application
}

func (a *AuthServiceServer) VerifyUser(ctx context.Context, t *pb.Token) (*pb.AuthResponse, error) {
	token := t.Token

	if token == "" {
		return &pb.AuthResponse{
			Message: "token should be provided in a header",
			User:    nil,
		}, nil
	}

	tParts := strings.Split(token, " ")

	if len(tParts) != 2 {
		return &pb.AuthResponse{
			Message: "token should containt two parts",
			User:    nil,
		}, nil
	}

	if tParts[0] != "Bearer" {
		return &pb.AuthResponse{
			Message: "first part of the auth header should be \"Bearer\"",
			User:    nil,
		}, nil
	}

	login, err := a.app.parseToken(tParts[1], true)

	if err != nil {
		return &pb.AuthResponse{
			Message: err.Error(),
			User:    nil,
		}, nil
	}

	user, err := a.app.UserUsecases.GetUserByLogin(login)

	if err != nil {
		return &pb.AuthResponse{
			Message: err.Error(),
			User:    nil,
		}, nil
	}

	return &pb.AuthResponse{
		Message: "user verified",
		User: &pb.User{
			Id:        user.Id,
			Iin:       user.Iin,
			Login:     user.Login,
			Role:      nil,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (app *application) StartGrpcServer() error {
	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", app.config.Server.Grpc.Port))

	if err != nil {
		app.Logger.Printf("Error starting grpc server %v\n", err)
		return err
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	server := &AuthServiceServer{
		app: app,
	}

	pb.RegisterAuthServiceServer(grpcServer, server)

	app.Logger.Println("starting grpc server ... ")
	err = grpcServer.Serve(listener)
	if err != nil {
		app.Logger.Printf("error starting grpc server %v ", err)
	}
	return nil
}
