package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	xds "github.com/envoyproxy/go-control-plane/pkg/server/v2"
	"google.golang.org/grpc"
)

var (
	Version       int
	snapshotCache cache.SnapshotCache
	namespace     string
	config        *rest.Config
	err           error
)

func main() {

	namespace = os.Getenv("NAMESPACE")
	if len(namespace) == 0 {
		namespace = "default"
	}

	modeFlag := flag.String("mode", "local", "If local, use specified kubeconfig. If cluster, use in-pod config")
	kubeConfig := flag.String("kubeConfig", "./.kube/config", "path to kubeconfig")
	flag.Parse()

	var watcher Watcher
	if *modeFlag == "local" {
		log.Println("Using config: ", *kubeConfig)
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			log.Fatalln(err)
		}
	} else if *modeFlag == "cluster" {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalln(err)
		}

	}

	Version = 0
	snapshotCache = cache.NewSnapshotCache(true, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, nil)
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":8081")

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	log.Println("Starting server...")

	watcher, err = NewPathfinderSnapshotWatcher(config, namespace)
	if err != nil {
		log.Fatalln("Error init watcher", err)
	}

	go watcher.Watch()

	grpcServer.Serve(lis)
}
