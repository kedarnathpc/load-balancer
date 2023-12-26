package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/kedarnathpc/load-balancer/internal/server"
)

// LoadBalancer holds the configuration for the load balancer.
type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []server.ServerInterface
}

// NewLoadBalancer creates a new LoadBalancer instance.
func NewLoadBalancer(port string, servers []server.ServerInterface) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}

// GetNextAvailableServer returns the next available server for load balancing.
func (lb *LoadBalancer) GetNextAvailableServer() server.ServerInterface {
	// Implementation of load balancing logic goes here
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

// ServerProxy forwards the request to the next available server.
func (lb *LoadBalancer) ServerProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("Forwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}
