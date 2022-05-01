# syntax=docker/dockerfile:1
FROM golang:1.18 as build

WORKDIR /go/src/app
ADD . ./
RUN go mod download
RUN go build -o /go/bin/app cmd/aqimdemo/*.go

FROM debian:stable-slim
RUN apt update && apt install -y tmux
COPY --from=build /go/src/app/downloads/nats-server /go/src/app/internal/evt/testdata/*.conf /go/bin/app /
RUN ls -l /
CMD ["/nats-server", "-c", "/nats-leaf.conf"]
