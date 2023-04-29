FROM golang:1.20 as builder
WORKDIR /senao
ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/senaoapp ./cmd

FROM alpine:3.16
WORKDIR /senao
RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl tzdata && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /senao/cmd/senaoapp /senao/senaoapp

EXPOSE 8080
