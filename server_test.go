package main

import (
	"context"
	"log"
	"testing"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"google.golang.org/grpc"
)

func TestDiscoverCluster(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := envoy_api_v2.NewClusterDiscoveryServiceClient(conn)
	req := envoy_api_v2.DiscoveryRequest{
		VersionInfo: "0",
		Node: &envoy_api_v2_core.Node{
			Id: "mesh",
		},
	}
	resp, err := client.FetchClusters(
		context.Background(),
		&req,
		grpc.EmptyCallOption{},
	)
	log.Println(resp)
}
