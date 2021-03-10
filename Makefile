KUBE_CONFIG ?= ./.kube/config
NAMESPACE ?= default
IMG ?= wylswz/trovu:0.1.0

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

docker-build:
	docker build --network=host -t $(IMG)  .

docker-push: docker-build
	docker push $(IMG)