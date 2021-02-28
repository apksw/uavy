DEV_TAG=dev
STG_TAG=stage
PROD_TAG=v0.0.1
IMAGE_NAME=uavy/auth


build:
	go build ./...

unit-test:
	@echo "FIX: execute only unit tests"
	go clean -testcache
	go test ./... -v

build-test:
	docker build -t $(IMAGE_NAME):$(DEV_TAG) -f Dockerfile.test .
	docker tag $(IMAGE_NAME):$(DEV_TAG) localhost:5000/$(IMAGE_NAME):$(DEV_TAG)

test:
	docker-compose down
	sudo rm -rf ./_data
	docker-compose up --force-recreate --build

