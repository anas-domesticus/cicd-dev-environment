FROM golang:1.22.1 as builder
WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o "example" main.go

FROM scratch
COPY --from=builder /src/example /usr/local/bin/example
ENTRYPOINT ["/usr/local/bin/example"]