# Variáveis
APP_NAME = tp
VERSION = 1.0.0

# Build local
build:
	go build -o bin/$(APP_NAME) cmd/server.go

# Testes
test:
	go test ./... -v

# Limpeza do build
clean:
	rm -rf bin/

# Build do Docker
docker-build:
	docker build -t $(APP_NAME):$(VERSION) .

# Rodar o container Docker
docker-run:
	docker run -d -p 8080:8080 --name $(APP_NAME) $(APP_NAME):$(VERSION)

# Parar e remover o container
docker-stop:
	docker stop $(APP_NAME) && docker rm $(APP_NAME)

