FROM golang:1.21.0
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0  go build -o /go-server
EXPOSE 8080
CMD ["/go-server"]
