include .env

stop_containers:
	@echo "Stopping other docker container"
	if [ $$(docker ps -q) ]; then \
		echo "Found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "No containers running..."; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5433:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine