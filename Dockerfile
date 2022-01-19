FROM golang:latest

COPY ./ ./
ENV GOPATH ""
ENV message "hi there"
RUN go build -o main
CMD ["./main"]
