FROM golang:1.24.5-alpine3.22

RUN mkdir /app
Add . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]