run:
	@go run main.go

up:
	@docker-compose up -d
	@docker ps -a

down:
	@docker-compose down
	@docker ps -a

clients:
	go run client/main.go --requests 100