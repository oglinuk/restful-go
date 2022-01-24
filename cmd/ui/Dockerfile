FROM golang:1.17 as build
WORKDIR /go/src/github.com/oglinuk/restful-go/ui
COPY . .
RUN go mod download && go mod verify
ENV CGO_ENABLED=0
RUN go build -ldflags="-w -s"

FROM scratch
COPY --from=build /go/src/github.com/oglinuk/restful-go/ui/static /static
COPY --from=build /go/src/github.com/oglinuk/restful-go/ui/templates /templates
COPY --from=build /go/src/github.com/oglinuk/restful-go/ui/ui ui
EXPOSE 9042
ENTRYPOINT ["/ui"]
