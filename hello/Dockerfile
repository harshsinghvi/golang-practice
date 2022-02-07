FROM golang:1.17 as builder
WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./...

FROM alpine
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
CMD ["app"]
