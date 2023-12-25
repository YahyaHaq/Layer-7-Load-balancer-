package main

import (
	loadbalancer "github.com/YahyaHaq/Layer-7-Load-balancer-/load_balancer"
)

func main() {
	lb := loadbalancer.NewLoadBalancer([]string{
		"http://localhost:5000",
		"http://localhost:5001",
	})

	port := 8080

	lb.Start(port)
}
