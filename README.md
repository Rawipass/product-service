# product-service

# migration cmd
migrate -database "postgres://postgres:postgres@localhost:5432/yourdbname?sslmode=disable" -path ./migrations up \n
migrate -database "postgres://postgres:postgres@localhost:5432/yourdbname?sslmode=disable" -path ./migrations down

# docker cmd
docker-compose up --build \n
docker-compose up -d

# run service
go run main.go
