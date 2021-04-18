FROM golang:alpine3.13 as builder

RUN apk update && apk add git
COPY . $GOPATH/src/github.com/lacazethomas/badTracker/
WORKDIR $GOPATH/src/github.com/lacazethomas/badTracker/
RUN go build -o /go/bin/badTracker

ENV ENVIRONMENT prod

FROM alpine:3.13.1
EXPOSE 8080
COPY --from=builder /go/bin/badTracker /bin/badTracker

ENTRYPOINT ["/bin/badTracker"]