FROM golang:1.16-alpine

WORKDIR /app

COPY presenter/go.mod ./main/go.mod
COPY ./data/go.mod ./data/go.mod
COPY ./domain/go.mod ./domain/go.mod

RUN cd ./presenter && go mod download

COPY presenter/*.go ./main
COPY ./data/*.go ./data
COPY ./domain/*.go ./domain

RUN cd ./presenter && go build -o /system-info-service

EXPOSE 8080

CMD [ "/system-info-service" ]