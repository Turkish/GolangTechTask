package api
//
//import (
//	"context"
//	"github.com/buffup/GolangTechTask/cmd"
//	"github.com/buffup/GolangTechTask/storage"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/test/bufconn"
//	"log"
//	"net"
//	"testing"
//)
//
//const bufSize = 1024 * 1024
//
//var lis *bufconn.Listener
//
//func init() {
//	lis = bufconn.Listen(bufSize)
//	s := grpc.NewServer()
//	RegisterVotingServiceServer(s, main.NewService(storage.R))
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}
//
//func bufDialer(context.Context, string) (net.Conn, error) {
//	return lis.Dial()
//}
//
//func TestCreateVoteable(t *testing.T) {
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	client := NewVotingServiceClient(conn)
//	resp, err := client.CreateVoteable(ctx, &CreateVoteableRequest{
//		Question:      "Did player deserve the RED card ?",
//		Answers:       []string{"YES", "NO"},
//	})
//	if err != nil {
//		t.Fatalf("CreateVoteable failed: %v", err)
//	}
//	log.Printf("Response: %+v", resp)
//	// Test for output here.
//}
