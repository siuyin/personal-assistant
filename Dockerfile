# syntax=docker/dockerfile:1
FROM golang:1.18 as build

WORKDIR /go/src/app
ADD . ./
RUN go mod download
RUN go build -o /go/bin/app cmd/aqimdemo/*.go

FROM debian:stable-slim
RUN apt update && apt install -y tmux
ADD https://github.com/nats-io/nats-server/releases/download/v2.8.1/nats-server-v2.8.1-linux-amd64.tar.gz /tmp/
RUN tar -C / -xf /tmp/nats-server-v2.8.1-linux-amd64.tar.gz
RUN mv /nats-server-v2.8.1-linux-amd64 /nats-server
COPY --from=build /go/src/app/internal/evt/testdata/*.conf /go/bin/app /
RUN ls -l /
#COPY --from=build /go/src/app/downloads/nats-server /go/src/app/internal/evt/testdata/*.conf /go/bin/app /
CMD ["/nats-server", "-c", "/nats-leaf.conf"]
