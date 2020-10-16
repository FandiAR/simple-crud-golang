FROM golang:1.12-alpine3.10 AS build-stage

LABEL maintainer="Fandi Agus Riyanto<fandi@bariqpratama.com>"
LABEL REPO="https://github.com/FandiAR/simple-crud-golang"

RUN apk add git && apk add --update make

WORKDIR /go/src/app

RUN go get github.com/gorilla/mux

COPY . .

RUN go build -o bin/simple-crud-golang

EXPOSE 8080

FROM alpine

LABEL maintainer="Fandi Agus Riyanto<fandi@bariqpratama.com>"
LABEL REPO="https://github.com/FandiAR/simple-crud-golang"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true 

COPY --from=build-stage /go/src/app/bin/simple-crud-golang /go/src/app/bin/

WORKDIR /go/src/app/bin

CMD ["/go/src/app/bin/simple-crud-golang"]