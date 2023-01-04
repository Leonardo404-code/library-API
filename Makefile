db_up:
	@docker compose -f ./config/docker/docker-compose.yml up -d --build

db_down:
	@docker compose -f ./config/docker/docker-compose.yml down