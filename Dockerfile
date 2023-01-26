FROM golang
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o ./mysuperAPI cmd/api/*.go
CMD ["/app/mysuperAPI"]



