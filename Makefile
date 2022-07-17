build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

test:
	go mod tidy
	go test -v .

docs:
	go get -v github.com/swaggo/swag/cmd/swag
	go get -u github.com/arsmn/fiber-swagger/v2
	go mod vendor -v
	swag init -g main.go --output docs --parseDependency --parseInternal

#################################################################
####################### When using Docker #######################
#################################################################

run-docker-image:
	docker build -f dockerfiles/golang/Dockerfile \
		--build-arg BUILD_WORKFLOW=production_code \
		-t flight_log_prod .
	docker run -p 8000:8000 flight_log_prod

test-docker-image:
	docker build -f dockerfiles/golang/Dockerfile \
		--build-arg BUILD_WORKFLOW=testing_code \
		-t flight_log_test .
	docker run flight_log_test
