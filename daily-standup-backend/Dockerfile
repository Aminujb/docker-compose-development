FROM golang:1.17-alpine

LABEL Jubril Aminu  <jubrilaminu304@gmail.com>

RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8000

RUN go mod download && go build -o main .

CMD ["/app/main"]
