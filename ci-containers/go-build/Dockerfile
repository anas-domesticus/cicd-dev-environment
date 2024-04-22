FROM golang:1.22.1 as builder
WORKDIR /src/ci-helper
COPY ci-helper/go.mod ./
RUN go mod download
COPY ci-helper/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ci-helper .

FROM golang:1.22.1
LABEL maintainer="Lewis Edginton"
RUN apt update && apt install -y git
COPY --from=builder /src/ci-helper/ci-helper /usr/local/bin/ci-helper
WORKDIR /app
