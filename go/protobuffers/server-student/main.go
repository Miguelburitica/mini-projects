package main

import (
	"log"
	"net"

	"github.com/Miguelburitica/mini-projects/go/protobuffers/database"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/server"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := server.NewStudentServer(repo)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}