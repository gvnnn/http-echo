package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort string = "8080"

func main() {
	var port string
	http.Handle("/echo", http.HandlerFunc(echo))

	if port = os.Getenv("ECHO_LISTEN_PORT"); port == "" {
		port = defaultPort
	}

	fmt.Printf("Listening on %s\n", ":" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", r.Host)

	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
