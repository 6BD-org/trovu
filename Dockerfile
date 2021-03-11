FROM golang:1.15 as builder

ENV https_proxy 10.114.114.1:1091

WORKDIR /workspace

ADD . /workspace

RUN make build

FROM ubuntu:20.04

WORKDIR /

COPY --from=builder /workspace/trovu /trovu

ENTRYPOINT ["/trovu", "-mode", "cluster"]