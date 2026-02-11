# syntax=docker/dockerfile:1.7

FROM golang:1.23-bookworm AS build
WORKDIR /app/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/tcc .

FROM debian:bookworm-slim AS runtime
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash parallel ca-certificates coreutils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app/src
COPY --from=build /out/tcc /app/src/tcc
COPY run_all.sh /app/src/run_all.sh
RUN chmod +x /app/src/tcc /app/src/run_all.sh

ENTRYPOINT ["/app/src/tcc"]
CMD ["--help"]
