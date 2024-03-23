FROM golang:1.22.1-alpine

WORKDIR /app

COPY . .

RUN go build -o payment ./cmd/server

EXPOSE 8001

CMD [ "./payment"]