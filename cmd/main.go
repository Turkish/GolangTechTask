package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/buffup/GolangTechTask/api"
	"github.com/buffup/GolangTechTask/storage"
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

	_, err = s.CreateVoteable(&api.Voteable{Question: "Did player deserve red card?", Answers: []string{"YES", "NO"}})

	//
	//
	//lis, err := net.Listen("tcp", ":9000")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//grpcServer := grpc.NewServer()
	//RegisterVotingServiceServer(grpcServer, NewService(s))
	//if err := grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
}

