FROM golang:latest AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/kikils/golang-todo
COPY ./ ./
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/kikils/golang-todo /app
COPY --from=builder /go/src/github.com/kikils/golang-todo/firebase-key.json /.key/firebase-key.json

CMD /app/main $PORT