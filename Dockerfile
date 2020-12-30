FROM golang:1.13 as builder

WORKDIR /workspace

ADD . /workspace

RUN make build

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /workspace/trovu .
USER nonroot:nonroot

ENTRYPOINT ["./trovu"]