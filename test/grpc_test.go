package test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/asnur/go-grpc-learn/student"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func getDataStudent(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if email == "" {
		resp, err := client.FindAllStudent(ctx, new(empty.Empty))
		if err != nil {
			log.Fatalf("could not find student: %v", err)
		}
		log.Printf("AllStudent: %v", resp)
	} else {
		student := &pb.Student{
			Email: email,
		}
		resp, err := client.FindStudentByEmail(ctx, student)
		if err != nil {
			log.Fatalf("could not find student with email %s: %v", email, err)
		}
		log.Printf("student: %s", resp.Name)
	}

}

func DialServer() grpc.ClientConnInterface {
	var opt []grpc.DialOption

	opt = append(opt, grpc.WithInsecure())
	opt = append(opt, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:5000", opt...)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return conn
}

func TestDial(t *testing.T) {
	conn := DialServer()

	client := pb.NewDataStudentClient(conn)
	getDataStudent(client, "")
}
