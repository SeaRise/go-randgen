FROM golang:1.16 as builder
RUN go get -u github.com/jteeuwen/go-bindata/...
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the source
COPY . .

# Build
RUN GOOS=linux GOARCH=amd64 make bin

FROM pingcap/alpine-glibc:3.10

RUN apk add --no-cache bash curl mysql-client

COPY --from=builder /workspace/go-randgen /go-randgen
COPY --from=builder /workspace/resource /resource
COPY --from=builder /workspace/cases /cases
COPY --from=builder /workspace/run.sh /run.sh
