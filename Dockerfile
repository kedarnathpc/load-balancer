FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go install ./...

WORKDIR /go/src/app/cmd

RUN go mod download

RUN go build -o main .

EXPOSE 8080

# Command to run the executable
CMD ["./main"]
