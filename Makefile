run:
	@go run main.go

docker:
	@docker-compose up -d

down:
	@docker-compose down

success:
	@curl http://127.0.0.1:8080/

fail:
	@curl http://127.0.0.1:8080/fail

metrics:
	@curl http://127.0.0.1:8080/metrics