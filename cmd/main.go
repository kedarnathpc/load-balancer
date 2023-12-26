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
		server.NewServer("https://bing.com/"),
		server.NewServer("https://example.com/"),
		server.NewServer("https://wikipedia.org/"),
		server.NewServer("https://reddit.com/"),
		server.NewServer("https://stackoverflow.com/"),
		server.NewServer("https://netflix.com/"),
		server.NewServer("https://amazon.com/"),
		server.NewServer("https://apple.com/"),
		server.NewServer("https://spotify.com/"),
		server.NewServer("https://microsoft.com/"),
		server.NewServer("https://instagram.com/"),
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
