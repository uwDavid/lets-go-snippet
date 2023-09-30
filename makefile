server: 
	go run cmd/web/!(*_test).go

d-up:
	docker compose up -d

d-down: 
	docker compose down

test:
	go test -failfast -v ./cmd/web

test-integration:
	go test -v ./pkg/models/mysql

# skip long-running tests
test-short:
	go test -v -short ./...

# show test coverage
test-cover: 
	go test -cover ./...

# get coverage profile
test-profile: 
	go test -coverprofile=/tmp/profile.out ./...

