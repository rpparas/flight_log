#!make
SHELL := /bin/bash

build:
	docker build -t postgres .

run-local:
	docker rm postgres || true
	docker run -d -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgres postgres

run-ci:
	docker rm postgres || true
	docker run -d -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -p 5432:5432 --name postgres postgres
