build:
	docker build -t postgres .

run:
	docker run \
		-e POSTGRES_USER=myusername \
		-e POSTGRES_PASSWORD=mypassword \
		-p 5432:5432 \
		-v /data:/var/lib/postgresql/data \
		postgres
