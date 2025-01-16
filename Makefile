runbackend:
	go run backend/main.go

postgres:
	docker exec -it url-shortener-postgres psql

postgresinit:
	docker run --name url-shortener-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17-alpine

.PHONY: runbackend postgres postgresinit