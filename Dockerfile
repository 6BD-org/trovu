FROM golang:1.13 as builder

ENV https_proxy 10.114.114.1:1091

WORKDIR /workspace

ADD . /workspace

RUN make build

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /workspace/trovu .
USER nonroot:nonroot

ENTRYPOINT ["./trovu", "-mode", "cluster"]