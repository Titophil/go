
---

### 5. `Dockerfile`
```dockerfile
# Multi-stage build
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o hello

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/hello .
EXPOSE 8080
CMD ["./hello"]
