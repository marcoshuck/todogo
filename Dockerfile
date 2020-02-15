# Builder --
FROM golang:latest as builder
LABEL maintainer="Marcos Huck <marcos@huck.com.ar>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Runner --
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 3000
CMD ["./main"]
