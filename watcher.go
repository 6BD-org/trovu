package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	v1 "github.com/6BD-org/pathfinder/api/v1"
	client "github.com/6BD-org/pathfinder/client"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"k8s.io/client-go/rest"
)

type Watcher interface {
	Watch()
}

type PathfinderSnapshotWatcher struct {
	client    client.XMClient
	namespace string
}

func (psw PathfinderSnapshotWatcher) Watch() {
	last := v1.PathFinderList{}
	log.Println("Using namespace: ", psw.namespace)
	for {
		time.Sleep(5 * time.Second)
		new := v1.PathFinderList{}
		psw.client.PathFinderV1(namespace).List(context.TODO(), &new, client.PathFinderListOption{})
		if !reflect.DeepEqual(last, new) {
			log.Println("Service changed, updating shapshot")
			last = new
			clusters := make([]types.Resource, 0)
			for _, v := range new.Items {
				for _, entry := range v.Status.ServiceEntries {
					hosts := make([]*envoy_api_v2_core.Address, 1)
					addr := envoy_api_v2_core.Address{
						Address: &envoy_api_v2_core.Address_SocketAddress{
							SocketAddress: &envoy_api_v2_core.SocketAddress{
								Protocol: 0,
								Address:  entry.ServiceHost,
							},
						},
					}

					hosts[0] = &addr
					clusters = append(clusters, &envoy_api_v2.Cluster{
						Name: entry.ServiceName,
						ClusterDiscoveryType: &envoy_api_v2.Cluster_Type{
							Type: envoy_api_v2.Cluster_STRICT_DNS,
						},
						Hosts: hosts,
					})
				}
			}
			log.Println("Generating snapshot for cluster: ", clusters)
			snapshot := cache.NewSnapshot(fmt.Sprintf("%v", Version), nil, clusters, nil, nil, nil, nil)
			Version++
			snapshotCache.SetSnapshot("mesh", snapshot)

		} else {
			log.Println("Service Unchanged")
		}
	}
}

func NewPathfinderSnapshotWatcher(config *rest.Config, namespace string) (Watcher, error) {
	client, err := client.New(config)
	if err != nil {
		return PathfinderSnapshotWatcher{}, err
	}
	return PathfinderSnapshotWatcher{
		client:    client,
		namespace: namespace,
	}, nil
}
