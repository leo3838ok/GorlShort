.PHONY: dev
dev:
	cp .env.docker .env
	go mod vendor
	docker-compose up --build -d
