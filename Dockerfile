FROM golang:alpine

WORKDIR /go/src/app

RUN apk add git

RUN go get github.com/gorilla/mux

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]