APP_ENV = develop
MAIN_CONTAINER = app

COMPOSE = docker-compose -f docker/docker-compose.yml -p golang

build: service_build

service_build:
	$(COMPOSE) build

up:
	$(COMPOSE) up ${ARG}

down:
	$(COMPOSE) down

start: 
	make up ARG=-d

stop:
	$(COMPOSE) stop

restart: stop start

reload:
	$(COMPOSE) stop  $(MAIN_CONTAINER)
	$(COMPOSE) up -d $(MAIN_CONTAINER)

sh:
	$(COMPOSE) run $(MAIN_CONTAINER) sh

clean:
	$(COMPOSE) down -v --rmi all