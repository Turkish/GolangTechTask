package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/buffup/GolangTechTask/api"
	"github.com/buffup/GolangTechTask/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c := &aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("DUMMYIDEXAMPLE", "DUMMYEXAMPLEKEY", ""),
		Endpoint: aws.String("http://localhost:8000"),
	}
	s, err := storage.New(c)
	if err != nil {
		panic(err)
	}
	svc := api.NewVotingService(*s)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterVotingServiceServer(grpcServer, svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

