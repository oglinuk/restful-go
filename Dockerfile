FROM golang:1.17
WORKDIR /go/src/github.com/oglinuk/restful-go
COPY . .
RUN go mod download && go mod verify
WORKDIR ./cmd/api
RUN go build -ldflags="-w -s"
EXPOSE 9001
ENTRYPOINT ["./api"]
