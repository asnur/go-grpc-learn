package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	pb "github.com/asnur/go-grpc-learn/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	students []*pb.Student
}

func (s *dataStudentServer) FindStudentByEmail(ctx context.Context, in *pb.Student) (*pb.Student, error) {
	for _, student := range s.students {
		if student.Email == in.Email {
			return student, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "student with email %s not found", in.Email)
}

func (s *dataStudentServer) loadData() {
	data, err := os.ReadFile("data/student.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &s.students)
	if err != nil {
		log.Println(err)
	}

}

func newServer() *dataStudentServer {
	s := &dataStudentServer{}
	s.loadData()

	return s
}

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
