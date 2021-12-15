FROM golang:1.16-alpine

WORKDIR /app

COPY /hub/go.mod ./
COPY /hub/go.sum ./

RUN go mod download

COPY /hub/*.go ./

RUN go build -o /main

EXPOSE 3000

CMD [ "/main" ]