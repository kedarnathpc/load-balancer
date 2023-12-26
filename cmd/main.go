package main

import (
	"fmt"
	"net/http"

	"github.com/kedarnathpc/load-balancer/internal/loadbalancer"
	"github.com/kedarnathpc/load-balancer/internal/server"
)

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	servers := []server.ServerInterface{
		server.NewServer("https://google.com/"),
		server.NewServer("https://www.bing.com/"),
		server.NewServer("https://github.com/"),
	}

	lb := loadbalancer.NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.ServerProxy(rw, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at localhost:%s\n", lb.Port)

	http.ListenAndServe(":"+lb.Port, nil)
}
