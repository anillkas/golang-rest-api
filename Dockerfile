FROM golang:latest
COPY . /go/src/github.com/anillkas/golang-rest-api

WORKDIR /go/src/github.com/anillkas/golang-rest-api
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/anillkas/golang-rest-api .
CMD ["./app"]
