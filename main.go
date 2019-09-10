package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
)

var (
	port = flag.Int("port", 8080, "port")
	path = flag.String("path", "/", "health check path")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort("", strconv.Itoa(*port))

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		return
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}
