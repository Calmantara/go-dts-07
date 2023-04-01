build:
	make -C ./postgres build-image
	make -C ./nginx build-image
	make -C ./app/go-dts-user build-image
	make -C ./redis build-image
up:
	docker-compose --compatibility -f docker/compose.yaml up --build -d
up-stack:
	docker stack deploy --compose-file docker/stack.yaml --orchestrator swarm dts-07
down:
	docker-compose stop -t 1
down-stack: 
	docker stack rm dts-07
