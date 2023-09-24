server: 
	go run cmd/web/*

d-up:
	docker compose up -d

d-down: 
	docker compose down