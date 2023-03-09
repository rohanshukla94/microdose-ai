project_name = microdose
image_name = microdose:latest

run-local:
	go run .

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

up:
	make up-silent
	make shell

build:
	docker build -t $(image_name) .

build-no-cache:
	docker build --no-cache -t $(image_name) .

up-silent:
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./microdose

up-silent-prefork:
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./microdose -prod

delete-container-if-exist:
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell:
	docker exec -it $(project_name) /bin/sh

stop:
	docker stop $(project_name)

start:
	docker start $(project_name)
