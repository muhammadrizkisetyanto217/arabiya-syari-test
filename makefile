# Gunakan image Go resmi
FROM golang:1.23.2 AS builder

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependency
RUN go mod tidy

# Build aplikasi (sesuaikan path main.go di dalam cmd)
RUN go build -o main ./cmd

# Gunakan image minimal untuk menjalankan aplikasi
FROM alpine:latest

# Set working directory di dalam container
WORKDIR /root/

# Copy binary dari builder
COPY --from=builder /app/main .

# Jalankan aplikasi
CMD ["./main"]
