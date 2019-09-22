FROM golang:1.13 AS builder
ADD . /src
RUN go get -u github.com/gorilla/mux && cd /src && CGO_ENABLED=0 GOOS=linux go build -a -o webapp /src/cmd/api/api.go

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
COPY --from=builder /src/webapp /app/
ENTRYPOINT [ "./webapp" ]

