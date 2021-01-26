FROM golang:1.14.2-alpine as build_base

RUN apk add --no-cache git build-base

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/ecom-api .

# Runtime container
FROM alpine:3.12
RUN apk add ca-certificates postgresql-client

RUN addgroup -g 1000 ecom_group

RUN adduser -D -s /bin/sh -u 1000 -G ecom_group ecom_user

COPY --from=build_base /go/src/app/out/ecom-api /app/ecom-api

EXPOSE 8080

RUN chown ecom_user:ecom_group /app/ecom-api

USER ecom_user

CMD ["/app/ecom-api"]