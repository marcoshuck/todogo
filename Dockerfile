# Builder --
FROM golang:latest as builder
LABEL maintainer="Marcos Huck <marcos@huck.com.ar>"
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Runner --
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/app .
EXPOSE 3000
CMD ["./main"]