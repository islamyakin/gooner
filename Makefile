# Makefile

# Nama aplikasi dan image Docker
APP_NAME := gooner
DOCKER_IMAGE := gooner:v1

# Direktori kerja saat menjalankan perintah make
PWD := $(shell pwd)

# Perintah untuk membangun aplikasi Go
build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME) .

# Perintah untuk membangun image Docker
build-docker: build
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Perintah untuk menjalankan kontainer Docker
run: build-docker
	@echo "Running Docker container..."
	docker run -d -p 9090:9090 -v /var/run/docker.sock:/var/run/docker.sock $(DOCKER_IMAGE)

# Membersihkan file binary dan image Docker
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
	docker rmi -f $(DOCKER_IMAGE)

# Menjalankan make build dan make run
all: build run
