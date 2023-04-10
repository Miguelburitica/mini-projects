package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// originServerUrl, err := url.Parse(envvar.Get("ORIGIN_SERVER_URL"))
	originServerUrl, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("Invalid origin server URL")
	}
	
	reverseProxyHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())
		
		req.Host = originServerUrl.Host
		req.URL.Host = originServerUrl.Host
		req.URL.Scheme = originServerUrl.Scheme
		req.RequestURI = ""
		
		req.Header.Add("X-Forwarded-For", req.RemoteAddr)
		req.Header.Add("X-Forwarded-Proto", req.URL.Scheme)
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("Forwarded", "for="+req.RemoteAddr+";host="+req.Host+";proto="+req.URL.Scheme)

		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		rw.Header().Set("Server", "gxh - Reverse Proxy Server")
		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, originServerResponse.Body)

	},)

	log.Fatal(http.ListenAndServe(":8080", reverseProxyHandler))
}