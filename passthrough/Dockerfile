FROM golang:1.22-alpine

# This avoids the "go.mod exists but should not" error
WORKDIR /go/delivery

COPY . .

RUN cd src/ && /usr/local/go/bin/go build -o "../bin/passthrough"

# HTTP port
EXPOSE 7500
# ICE port 
EXPOSE 40000/udp

ENTRYPOINT ["/go/delivery/bin/passthrough", "-debug"]

