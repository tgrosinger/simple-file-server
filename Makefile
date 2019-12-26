
.PHONY: simple-file-server
simple-file-server:
	CGO_ENABLED=0 go build -o bin/simple-file-server cmd/server/main.go

.PHONY: docker-container
docker-container:
	docker build -t tgrosinger/simple-file-server$(DOCKER_IMAGE_TAG) .
