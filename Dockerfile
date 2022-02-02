FROM golang:1.17-alpine as BUILD

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /app
RUN go build -o ./mirrorz-shim

FROM alpine:3

RUN apk add ca-certificates

WORKDIR /app
COPY --from=BUILD /app/mirrorz-shim /app/mirrorz-shim
COPY configs ./configs

EXPOSE 1323

CMD [ "./mirrorz-shim" ]
