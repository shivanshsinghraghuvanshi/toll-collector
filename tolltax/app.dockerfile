FROM golang:1.15-alpine3.13 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/ssinghraghuvanshi/toll-collector
COPY go.mod go.sum ./
COPY vendor vendor
COPY tolltax tolltax
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./tolltax/cmd/tolltax

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
