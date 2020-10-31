FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/temp
COPY . .
RUN go build ./src/cmd/main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/temp /app

CMD /app/main $PORT