FROM golang:1.16-alpine

WORKDIR /app

COPY ./main/go.mod ./main/go.mod
COPY ./data/go.mod ./data/go.mod
COPY ./domain/go.mod ./domain/go.mod

RUN cd ./main && go mod download

COPY ./main/*.go ./main
COPY ./data/*.go ./data
COPY ./domain/*.go ./domain

RUN cd ./main && go build -o /system-info-service

EXPOSE 8080

CMD [ "/system-info-service" ]