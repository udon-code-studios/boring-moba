# file: Dockerfile
#

FROM golang:1.13.8

ADD src src

RUN go get github.com/gorilla/websocket 
RUN go build src/server.go

CMD ["./server"]

EXPOSE 8080

#
# end of file
