FROM golang:1.15.2-alpine as build_base

RUN apk add --no-cache git build-base

WORKDIR /go/src/app/

COPY go.mod .
COPY go.sum .

COPY common-lib-image ../common-lib

RUN go mod download

COPY . .

RUN go build -o ./out/ecom-api .

CMD ["/app/ecom-api"]