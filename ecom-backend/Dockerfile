FROM golang:1.15.2-alpine as build_base

RUN apk add --no-cache git build-base

WORKDIR /go/src/app/

COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ecom-backend .

FROM alpine:3.12
RUN apk add ca-certificates postgresql-client

RUN addgroup -g 1000 ecom_group

RUN adduser -D -s /bin/sh -u 1000 -G ecom_group ecom_admin

COPY --from=build_base /go/src/app/out/ecom-backend /app/ecom-backend

EXPOSE 8080

RUN chown ecom_admin:ecom_group /app/ecom-backend

USER ecom_admin


CMD ["/app/ecom-backend"]