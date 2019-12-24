package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	version := "0.0.3"
	port := "8081"

	_ver, present := os.LookupEnv("VERSION")
	if present {
		version = _ver
	}

	_port, present := os.LookupEnv("PORT")
	if present {
		port = _port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "v%s - Hello from %s\n", version, hostname)
	})

	log.Printf("Starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
