ENV_FILE := .env
ENV = $(shell grep -v "^\#" $(ENV_FILE))
include .env

DOCKER_COMPOSE := docker compose

.PHONY: run frontend compose/up compose/down compose/logs

run: compose/up frontend

frontend:
	make -C frontend dev

compose/up:
	$(DOCKER_COMPOSE) up -d

compose/down:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans --rmi local

compose/logs:
	$(DOCKER_COMPOSE) logs -f


