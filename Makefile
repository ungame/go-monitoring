run:
	@go run main.go

up:
	@docker-compose up -d
	@docker ps -a

down:
	@docker-compose down
	@docker ps -a

clear:
	@docker rm -f go-monitoring
	@docker rm -f prometheus
	@docker rm -f grafana
	@docker ps -a
	@docker rmi -f go-monitoring_webapi
	@docker images -a

clients:
	go run client/main.go --requests 100