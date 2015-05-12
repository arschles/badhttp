package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	reqName = "request_name"
	delay   = "delay"
	code    = "code"
)

func main() {
	port := flag.Int("port", 8080, "the port to run on")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("/admin/{%s}/delay/{%s}", reqName, delay), adminDelay).Methods("PATCH")
	router.HandleFunc(fmt.Sprintf("/admin/{%s}/code/{%s}", reqName, code), adminCode).Methods("PATCH")
	router.HandleFunc(fmt.Sprintf("/{%s}", reqName), handler)

	hostStr := fmt.Sprintf(":%d", *port)
	fmt.Println("listening on", hostStr)
	log.Fatal(http.ListenAndServe(hostStr, router))
}
