build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

test:
	go mod tidy
	go test -timeout 30s -v .

test-csv:
	go test -timeout 30s -run ^TestPostFlightsCsv

go-docs:
	go get -v github.com/swaggo/swag/cmd/swag
	go get -u github.com/arsmn/fiber-swagger/v2
	go mod vendor -v
	swag init --parseDependency

dump-db:
	pg_dump --file db/flight_log.sql --host localhost \
		--port 5432 --username myusername --password \
		--no-blobs postgres

restore-db:
	cat db/flight_log.sql | docker exec -i postgres psql -U myusername

#################################################################
####################### When using Docker #######################
#################################################################

run-compose:
	docker network create flight_network || true
	docker-compose up --build
