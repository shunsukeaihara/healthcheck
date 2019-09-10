package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
)

var (
	port = flag.Int("poer", 8080, "int flag")
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
