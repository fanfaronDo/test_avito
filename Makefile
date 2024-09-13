run:
	go run ./cmd/app/main.go

migrateup:
	migrate -path ./migrations -database postgres://postgres:postgres@172.18.1.2:5432/avito?sslmode=disable -verbose up

migratedown:
	migrate -path ./migrations -database postgres://postgres:postgres@172.18.1.2:5432/avito?sslmode=disable -verbose down
