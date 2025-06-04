# Build the backend service
FROM golang:1.24.2-bullseye AS builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY ray-operator/ ray-operator/
COPY kwok/ kwok/

WORKDIR /workspace/kwok


# Build
USER root
WORKDIR /workspace/kwok
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o kuberay-kwok cmd/main.go

FROM scratch
WORKDIR /workspace
COPY --from=builder /workspace/kwok/kuberay-kwok kwok/
USER 65532:65532

ENTRYPOINT ["/workspace/kwok/kuberay-kwok"]
