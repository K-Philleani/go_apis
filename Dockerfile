FROM alpine:latest

WORKDIR /go/app

COPY ..

EXPOSE 9090

ENTRYPOINT ["/go/app/main"]