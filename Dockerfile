# Gunakan image golang:alpine sebagai base image
FROM golang

# Set working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum ke dalam container
COPY . .

# Download dependensi Go
RUN go mod download

# Salin kode sumber aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o crm-golang

# Menjalankan aplikasi saat container dijalankan
CMD ["./crm-golang"]
