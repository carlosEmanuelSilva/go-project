# Usa uma imagem base do Go
FROM golang:1.23 as builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos para o container
COPY . .

# Instala as dependências
RUN go mod download

# Compila o projeto
RUN go build -o server cmd/server.go

# Segunda etapa para criar a imagem final
FROM debian:latest
WORKDIR /root/

# Copia o binário do builder
COPY --from=builder /app/server .

# Exponha a porta
EXPOSE 8080

# Comando para rodar o servidor
CMD ["./server"]

