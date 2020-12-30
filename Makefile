KUBE_CONFIG ?= ./.kube/config
NAMESPACE ?= default

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

build: vet fmt
	go build -a -o trovu

run: build
	export NAMESPACE=${NAMESPACE} &&./trovu --mode=local --kubeConfig=${KUBE_CONFIG}

exec:
	export NAMESPACE=${NAMESPACE} &&./trovu --mode=local --kubeConfig=${KUBE_CONFIG}