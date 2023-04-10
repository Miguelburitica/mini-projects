package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "site to scan")

func main() {
	// 1, 2, 3, ..., 65535
	// sitio:1, sitio:2, sitio:3, ..., sitio:65535
	// 1 -> open, 2 -> closed
	
	flag.Parse()
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				// port is closed or filtered.
				return
			}

			conn.Close()
			fmt.Printf("port %d is open\n", port)
		}(i)
	}
	wg.Wait()
}