FROM golang:1.15.3 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/line_notify
COPY . .
RUN go build ./src/cmd/main.go

# runtime image
FROM alpine
RUN apk update
RUN apk add curl
COPY --from=builder /go/src/github.com/line_notify /app

CMD /app/main $PORT