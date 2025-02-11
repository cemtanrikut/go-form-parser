# Golang resmi imajını kullan
FROM golang:1.20

# Çalışma dizinini belirle
WORKDIR /app

# Bağımlılıkları yükle
COPY go.mod go.sum ./
RUN go mod download

# Proje dosyalarını kopyala
COPY . .

# Uygulamayı çalıştır
CMD ["go", "run", "main.go"]
