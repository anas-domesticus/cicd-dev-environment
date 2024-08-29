FROM golang:1.22.5 as builder
WORKDIR /src/ci-helper
COPY ci-helper/go.mod ./
RUN go mod download
COPY ci-helper/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ci-helper .

FROM golang:1.22.5
LABEL maintainer="Lewis Edginton"
RUN apt update && apt install -y git
RUN wget https://github.com/dagger/dagger/releases/download/v0.12.5/dagger_v0.12.5_linux_amd64.tar.gz && \
    tar -xvf dagger_v0.12.5_linux_amd64.tar.gz && \
    mv dagger /usr/local/bin/dagger && \
    rm LICENSE dagger_v0.12.5_linux_amd64.tar.gz
COPY --from=builder /src/ci-helper/ci-helper /usr/local/bin/ci-helper
WORKDIR /app
