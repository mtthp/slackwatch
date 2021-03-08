FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app/cmd
RUN go build -o main .
ENTRYPOINT ["/app/cmd/main"]
