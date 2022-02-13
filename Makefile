up:
	docker-compose up --build -d

down:
	docker-compose down

tail:
	docker logs -f $$(docker ps -f 'name=fantom-sales-bot-bot-1' --format='{{.ID}}')
