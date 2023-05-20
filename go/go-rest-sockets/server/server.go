package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Miguelburitica/go-rest-server/database"
	"github.com/Miguelburitica/go-rest-server/repository"
	"github.com/Miguelburitica/go-rest-server/websocket"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseURL string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub 	 *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DataBaseURL == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub: 	  websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func (b *Broker) Start(binder func (s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.AllowAll().Handler(b.router)

	repo, err := database.NewPostgresRepository(b.config.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)
	log.Println("Starting server on port " + b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
