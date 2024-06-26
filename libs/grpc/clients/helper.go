package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func createClientAddr(isProduction bool, service string, port int) string {
	ip := fmt.Sprintf("dns:///%s", service)
	if !isProduction {
		ip = "127.0.0.1"
	}

	return fmt.Sprintf("%s:%v", ip, port)
}

var defaultClientsOptions = []grpc.DialOption{
	grpc.WithTransportCredentials(insecure.NewCredentials()),
	grpc.WithDefaultServiceConfig(
		fmt.Sprintf(
			`{"loadBalancingPolicy":"%s", "lb_policy_name": "%s", "loadBalancingConfig": [{"%s":{}}]}`,
			roundrobin.Name,
			roundrobin.Name,
			roundrobin.Name,
		),
	),
	// grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
}
