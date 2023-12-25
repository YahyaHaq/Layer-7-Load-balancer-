package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type LoadBalancer struct {
	targets []*url.URL
	counter int
	Proxy   *httputil.ReverseProxy
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	targets := make([]*url.URL, 0, len((servers)))

	for _, server := range servers {
		u, err := url.Parse(server)
		if err != nil {
			fmt.Printf("Unable to parse url %s due to error %s", server, err)
			continue
		}

		targets = append(targets, u)
	}

	return &LoadBalancer{
		targets: targets,
		Proxy:   &httputil.ReverseProxy{Director: director},
	}
}

func (lb *LoadBalancer) Start(port int) {
	fmt.Printf("Load Balancer listening on :%d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), lb)
}

func director(r *http.Request) {
	r.Header.Set("X-Forwarded-For", r.RemoteAddr)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := lb.targets[lb.counter]
	lb.counter = (lb.counter + 1) % len(lb.targets)

	r.URL.Scheme = server.Scheme
	r.URL.Host = server.Host

	lb.Proxy.ServeHTTP(w, r)
}
