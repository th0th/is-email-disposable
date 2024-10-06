FROM golang:1.23.2

WORKDIR /usr/src/is-email-disposable

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN mkdir bin

RUN CGO_ENABLED=0 go build -a -o bin/rest-api cmd/restapi/*

FROM alpine:3

COPY --from=0 /usr/src/is-email-disposable/bin/rest-api /usr/local/bin/is-email-disposable-rest-api

ENTRYPOINT ["is-email-disposable-rest-api"]
