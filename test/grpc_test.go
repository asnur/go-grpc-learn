package test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/asnur/go-grpc-learn/student"
	"google.golang.org/grpc"
)

func getDataStudent(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	student := &pb.Student{
		Email: email,
	}

	resp, err := client.FindStudentByEmail(ctx, student)
	if err != nil {
		log.Fatalf("could not find student with email %s: %v", email, err)
	}

	log.Printf("student: %s", resp.Name)
}

func TestDial(t *testing.T) {
	var opt []grpc.DialOption

	opt = append(opt, grpc.WithInsecure())
	opt = append(opt, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:5000", opt...)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	getDataStudent(client, "asnurramdhani12@gmail.com")
}
