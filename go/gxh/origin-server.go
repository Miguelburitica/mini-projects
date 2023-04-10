package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	originServerHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Server", "Origin Server")

		fmt.Printf("[origin server] received request at: %s\n", time.Now())
		fmt.Printf("[origin server] origin addr: %s\n", req.Header.Get("X-Forwarded-For"))
		fmt.Printf("[origin server] origin host: %s\n", req.Header.Get("X-Forwarded-Host"))
		fmt.Printf("[origin server] origin proto: %s\n", req.Header.Get("X-Forwarded-Proto"))
		fmt.Printf("[origin server] origin forwardedInfo: %s\n", req.Header.Get("Forwarded"))

		if req.Method == http.MethodPost {
			rw.Write([]byte("Origin Server Response POST"))
		}

		if req.Method == http.MethodGet {
			time.Sleep(15 * time.Second)
			rw.Write([]byte("Origin Server Response GET"))
		}
	},)

	log.Fatal(http.ListenAndServe(":8081", originServerHandler))
}