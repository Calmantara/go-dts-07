build:
	make -C ./postgres build-image
	make -C ./nginx build-image
	make -C ./app/go-dts-user build-image
up:
	docker-compose --compatibility up --build -d
up-stack:
	docker stack deploy --compose-file docker-compose.yaml --orchestrator swarm dts-07
down:
	docker-compose stop -t 1