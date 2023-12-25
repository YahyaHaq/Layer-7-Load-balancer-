package loadbalancer_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	loadbalancer "github.com/YahyaHaq/Layer-7-Load-balancer-/load_balancer"
	"github.com/stretchr/testify/assert"
)

func Test_LoadBalancerIntegration(t *testing.T) {
	requestCount := map[string]int{}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the server's address from the URL
		// Write the server's address in the response
		key := r.Header.Get("X-Forwarded-For")

		requestCount[key] += 1
	})

	// create servers
	server1 := httptest.NewServer(handlerFunc)
	defer server1.Close()

	fmt.Println("Server1 ", server1.URL)

	server2 := httptest.NewServer(handlerFunc)
	defer server2.Close()

	fmt.Println("Server2 ", server2.URL)

	lb := loadbalancer.NewLoadBalancer([]string{server1.URL, server2.URL})

	port := 8000

	go lb.Start(port)

	for i := 0; i < 10; i++ {
		_, err := http.Get("http://localhost:8000/")
		assert.NoError(t, err)
	}

	fmt.Println(requestCount)

	assert.NotEqual(t, requestCount[server1.URL], 0)

	assert.Equal(t, requestCount[server1.URL], requestCount[server2.URL])
}
