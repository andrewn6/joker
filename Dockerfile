FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go mod download

RUN go build -o /api

EXPOSE 8000

CMD [ "/api" ]
