FROM golang:latest as build
WORKDIR $GOPATH/src/github.com/500deg/logging-microservices/users
COPY . .
RUN go version && go get -u -v golang.org/x/vgo
RUN vgo install ./...

FROM gcr.io/distroless/base
COPY --from=build /go/bin/users /
ENV PORT=8080
EXPOSE 8080
CMD ["/users"]