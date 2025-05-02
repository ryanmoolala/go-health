package main

import (
	"fmt"
	"log"
	"monitoring_probe/health"
	"net/http"
	"time"
)

var(
	endpoints []int
)

func main() {
	println("Starting monitoring probe on :8080")
	registerEndpoint(5002)
	registerEndpoint(5001)

	//run a check to see if all ports are available

	go routine_check() //Calls a check every few seconds

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server for performing health checks")
}

func registerEndpoint(endpoint int) {
	endpoints = append(endpoints, endpoint)
	fmt.Printf("Registered endpoint: %d\n", endpoint)
}


func routine_check() {
	for true {
		for _, endpoint := range endpoints {
			health.CheckHealth(endpoint)
		}
		time.Sleep(60 * time.Second)
	}
}


