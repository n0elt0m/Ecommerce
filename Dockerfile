FROM golang:latest
RUN mkdir /app
ADD . /app
ENV PORT 8000
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]