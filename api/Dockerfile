FROM golang:latest

WORKDIR /src
COPY . .
RUN mkdir config
RUN go build -o main
CMD ["./main"]