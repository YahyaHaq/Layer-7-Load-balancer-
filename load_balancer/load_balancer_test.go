package loadbalancer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLoadBalancer(t *testing.T) {
	servers := []string{"http:domain1.com/api", "https:domain2.com/api"}
	lb := NewLoadBalancer(servers)

	assert.Equal(t, len(servers), len(lb.targets))
}

func TestDirector(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://domain1.com/api", http.NoBody)
	req.RemoteAddr = "localhost"

	assert.NoError(t, err)

	director(req)

	assert.Equal(t, req.RemoteAddr, req.Header.Get("X-Forwarded-For"))
}
