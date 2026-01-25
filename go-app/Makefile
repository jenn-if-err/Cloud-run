
# Makefile for Go App Deployment to Cloud Run with Docker and gcloud (no Terraform)

PROJECT_ID=project-bc943afd-c417-4556-af8
REGION=asia-northeast1
APP_NAME=go-app
IMAGE=gcr.io/$(PROJECT_ID)/$(APP_NAME):latest

# Build the Docker image
build:
	docker build -t $(IMAGE) .

# Push the Docker image to Google Container Registry
push:
	docker push $(IMAGE)

# Authenticate Docker with Google Cloud
configure-docker:
	gcloud auth configure-docker


# Deploy to Cloud Run using gcloud
deploy:
	gcloud run deploy $(APP_NAME) \
	  --image $(IMAGE) \
	  --region $(REGION) \
	  --platform managed \
	  --allow-unauthenticated

# Full deploy: build, push, and deploy to Cloud Run
all: build push deploy

.PHONY: build push configure-docker deploy all
