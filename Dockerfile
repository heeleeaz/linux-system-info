FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./data/go.mod ./data/go.mod
COPY ./domain/go.mod ./domain/go.mod
COPY ./presenter/go.mod ./presenter/go.mod
RUN go mod download

COPY *.go ./
COPY ./data/*.go ./data
COPY ./domain/*.go ./domain
COPY ./presenter/*.go ./presenter

RUN go build -o /system-info-service

EXPOSE 8080

CMD [ "/system-info-service" ]