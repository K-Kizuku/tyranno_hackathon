ENV_FILE := .env
ENV = $(shell grep -v "^\#" $(ENV_FILE))
include .env

DOCKER_COMPOSE := docker compose

.PHONY: run backend frontend compose/up compose/down compose/logs gen-back clean gen gen-front

run: backend frontend

backend:
	go run backend/main.go   

frontend:
	make -C frontend dev

compose/up:
	$(DOCKER_COMPOSE) up -d

compose/down:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans --rmi local

compose/logs:
	$(DOCKER_COMPOSE) logs -f

gen-back:
	buf generate --template buf.gen.go.yaml 

gen-front:
	make -C frontend gen

gen: clean gen-back gen-front

gen-clean:
	rm -rf ./frontend/gen
	rm -rf ./backend/gen





