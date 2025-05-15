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
	// broker  *Broker
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
		// broker:  NewBroker("localhost:8024"),
	}
}

func (lb *LoadBalancer) Start(port int) {
	fmt.Printf("Load Balancer listening on :%d\n", port)

	// go lb.registerNewServers()

	http.ListenAndServe(fmt.Sprintf(":%d", port), lb)
}

// func (lb *LoadBalancer) registerNewServers() {
// 	pubsub := lb.broker.redis.Subscribe(context.Background(), "serverList")

// 	// Close the subscription when we are done.
// 	defer pubsub.Close()

// 	fmt.Println("waiting for servers to connect")
// 	for {
// 		msg, err := pubsub.ReceiveMessage(context.Background())
// 		if err != nil {
// 			panic(err)
// 		}

// 		serverInfo := &BrokerPayload{}

// 		fmt.Println(msg.Channel, msg.Payload)

// 		err = json.Unmarshal([]byte(msg.Payload), serverInfo)
// 		if err != nil {
// 			panic(err)
// 		}

// 		lb.addServer(serverInfo)

// 	}
// }

func (lb *LoadBalancer) addServer(serverInfo *BrokerPayload) {
	u, err := url.Parse(serverInfo.ServerAddress)
	if err != nil {
		fmt.Printf("Unable to parse url %s due to error %s", serverInfo.ServerAddress, err)
		return
	}

	lb.targets = append(lb.targets, u)

	fmt.Printf("added server %s", serverInfo.ServerAddress)
}

func director(r *http.Request) {
	r.Header.Set("X-Forwarded-For", r.RemoteAddr)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := lb.targets[lb.counter]
	lb.counter = (lb.counter + 1) % len(lb.targets)

	r.URL.Scheme = server.Scheme
	r.URL.Host = server.Host
	r.Host = server.Host

	lb.Proxy.ServeHTTP(w, r)
}
