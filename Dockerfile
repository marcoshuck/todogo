FROM golang:alpine
RUN mkdir -p /go/src/github.com/marcoshuck/todogo
COPY . /go/src/github.com/marcoshuck/todogo
WORKDIR /go/src/github.com/marcoshuck/todogo

RUN go build
CMD['todogo']

EXPOSE 3000