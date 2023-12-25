package loadbalancer_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	loadbalancer "github.com/YahyaHaq/Layer-7-Load-balancer-/load_balancer"
	"github.com/stretchr/testify/assert"
)

func Test_LoadBalancerIntegration(t *testing.T) {
	requestCount := map[string]int{}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount[r.Host] += 1
	})

	// create servers
	server1 := httptest.NewServer(handlerFunc)
	defer server1.Close()

	server2 := httptest.NewServer(handlerFunc)
	defer server2.Close()

	lb := loadbalancer.NewLoadBalancer([]string{server1.URL, server2.URL})

	port := 8000

	url := fmt.Sprintf("http://localhost:%d/", port)

	go lb.Start(port)

	for i := 0; i < 10; i++ {
		_, err := http.Get(url)
		assert.NoError(t, err)
	}

	assert.Equal(t, requestCount[get_key(server1.URL)], 5)

	assert.Equal(t, requestCount[get_key(server1.URL)], requestCount[get_key(server2.URL)])
}

// remove the leading scheme slashes from the URL to get it into same format as Host
func get_key(value string) string {
	return strings.Split(value, "://")[1]
}
