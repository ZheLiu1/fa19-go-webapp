FROM golang:1.13 AS builder
ADD . /src
RUN cd /src && make image

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
COPY --from=builder /src/bin/webapp /app/
ENTRYPOINT [ "./webapp" ]
