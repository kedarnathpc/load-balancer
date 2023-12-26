package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kedarnathpc/load-balancer/internal/utils"
)

type ServerInterface interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type simpleserver struct {
	address string
	proxy   *httputil.ReverseProxy
}

// NewServer creates a new simpleserver instance.
func NewServer(addr string) *simpleserver {
	serverURL, err := url.Parse(addr)
	utils.HandleErr(err)

	return &simpleserver{
		address: addr,
		proxy:   httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func (s *simpleserver) Address() string {
	return s.address
}

func (s *simpleserver) IsAlive() bool {
	return true
}

func (s *simpleserver) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
