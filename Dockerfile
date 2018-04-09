FROM golang:alpine
ADD . ./
RUN go build -o networkchecker
EXPOSE 8080
CMD ["./networkchecker"]