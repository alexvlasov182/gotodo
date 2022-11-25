up: #Containers start
	docker-compose up -d

down: #Stop
	docker-compose stop

ps:
	docker ps

connect_go:
	docker exec -it bash

connect_db:
	docker exec -it bash
