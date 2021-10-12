FROM golang:1.17
WORKDIR /go/src/github.com/oglinuk/restful-go/ui
COPY . .
RUN go mod download && go mod verify
RUN go build -ldflags="-w -s"
EXPOSE 9042
ENTRYPOINT ["./ui"]
