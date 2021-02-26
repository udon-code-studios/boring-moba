# file: Dockerfile
#

FROM golang:latest
ENV GOBIN /go/bin

WORKDIR /go/src/app
ADD src src

RUN go get github.com/gorilla/websocket 
RUN go build src/server.go

CMD ["./server"]

EXPOSE 8080

#
# end of file
