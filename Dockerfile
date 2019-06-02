# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10 AS build
#RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep
#RUN chmod +x $GOPATH/bin/dep
WORKDIR /go/src/github.com/opvizor/process-exporter
ADD . .
#RUN dep ensure

# Build the process-exporter command inside the container.
RUN make build

FROM trinitronx/jq

RUN apt-get update && apt-get install curl && rm -rf /var/lib/apt/lists/*

COPY --from=build /go/src/github.com/opvizor/process-exporter/process-exporter /bin/process-exporter

# Run the process-exporter command by default when the container starts.
ENTRYPOINT ["/bin/process-exporter"]

# Document that the service listens on port 9256.
EXPOSE 9584
