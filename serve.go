package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8000", "port to serve on")
	directory := flag.String("d", ".", "the directory of the static files to be hosted")
	flag.Parse()

	srv := http.FileServer(http.Dir(*directory))
	http.Handle("/", srv)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Printf("Access the server at http://127.0.0.1:%s/", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
