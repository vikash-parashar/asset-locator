# Makefile for Dockerized Golang/Gin application

# Environment variables for the Docker Compose file
export POSTGRES_USER=mydbuser
export POSTGRES_PASSWORD=mydbpassword
export POSTGRES_DB=mydbname

# Define the Docker image name
IMAGE_NAME=myapp
ACR_NAME=myacr  # Azure Container Registry name

.PHONY: build run db-create db-drop db-insert clean deploy

build:
	@docker build -t $(IMAGE_NAME) .

run:
	@docker-compose up

db-create:
	@docker-compose run --rm db psql -h db -U $(POSTGRES_USER) -d $(POSTGRES_DB) -a -f db/createTable.sql

db-drop:
	@docker-compose run --rm db psql -h db -U $(POSTGRES_USER) -d $(POSTGRES_DB) -a -f db/dropTable.sql

db-insert:
	@docker-compose run --rm db psql -h db -U $(POSTGRES_USER) -d $(POSTGRES_DB) -a -f db/insertData.sql

clean:
	@docker-compose down --volumes

# Azure deployment
deploy: build
	@az acr login --name $(ACR_NAME)
	@az acr build --registry $(ACR_NAME) --image $(IMAGE_NAME) .

# You can add more Azure-related deployment steps here

