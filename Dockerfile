# Example Dockerfile for a Go application (this is just a placeholder)
FROM scratch
COPY go-safe-conversion /
ENTRYPOINT ["/go-safe-conversion"]
