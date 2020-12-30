# Trovu

trovu is centralized service discoverer. It supports xDS standard and can be deployed with envoy service mesh

Still under design


## What does trovu do?
Trovu is a xDS standard cluster discovery service that lives in namespace. It watches couple of [pathfinders](https://github.com/6BD-org/pathfinder) and aggregate cluster endpoints. It wraps service entries into [envoy clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster.proto) in order to support service mesh deployment.

## What about endpoint discovery?
Kubernetes has already provided powerful DNS mechanism that's why trovu only provide [STRICT_DNS](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-strict-dns) in cluster discovery and leave load balance/ip selection to later stage

## Dependencies
- Pathfinder

## Deploy Trovu
not ready for deployment

## Run Trovu locally
Trovu requires access to kubernetes API (PathFinder API), so you may need to get ~/.kube/config ready. If you use [MediumKube](https://github.com/6BD-org/mediumkube) to deploy your test cluster, this file is automatically transferred to your preferred temp directory. Remember to unset proxy env before you use it, otherwize you might get `HTTP 403` in authorization.

To build and run

```sh
$ make run NAMESPACE={namespace where you install pathfinder}
```

To run executable without build
```sh
$ make exec NAMESPACE={namespace where you install pathfinder}
```