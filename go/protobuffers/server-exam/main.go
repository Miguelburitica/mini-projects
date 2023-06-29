package main

import (
	"log"
	"net"

	"github.com/Miguelburitica/mini-projects/go/protobuffers/database"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/exampb"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := server.NewExamServer(repo)

	s := grpc.NewServer()
	exampb.RegisterExamServiceServer(s, server)
	

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}