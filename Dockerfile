FROM golang:1.14.6

ADD . /go/src/gotodo
RUN go get github.com/lib/pq
RUN go install gotodo

ENTRYPOINT /go/bin/gotodo

EXPOSE 8080