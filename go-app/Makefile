PROJECT_ID=project-bc943afd-c417-4556-af8
REGION=asia-northeast1
REPO=my-repo
IMAGE=asia-northeast1-docker.pkg.dev/$(PROJECT_ID)/$(REPO)/go-app:latest

all: build push deploy

build:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)

deploy:
	gcloud run deploy go-app --image $(IMAGE) --region $(REGION) --allow-unauthenticated