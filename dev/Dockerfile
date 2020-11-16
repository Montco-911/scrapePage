FROM golang:1.15 as builder


WORKDIR /workspace
COPY go.mod go.mod

RUN go mod download


COPY . /workspace/
RUN go get -v -t -d ./...

RUN mkdir /pod-10g

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags timetzdata -a -o project main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /workspace/project .


COPY --from=builder --chown=65532:65532 /pod-10g /pod-10g
USER nonroot:nonroot


ENTRYPOINT ["/project"]
