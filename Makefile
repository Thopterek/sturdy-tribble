#
MAKEFLAGS += -s
IMAGE_NAME = sturdy-tribble
DEV_TAG = dev
PROD_TAG = prod

ifneq ($(wildcard .env),)
    include .env
    export
endif


# ==========================================
# Development Environment
# ==========================================

dev: dev-build connection-test dev-setup dev-run

dev-build:
	@echo "🔨 Building Docker image '$(IMAGE_NAME):$(DEV_TAG)'...\n"
	@if docker build \
		--target $(DEV_TAG) \
		--no-cache \
		-t $(IMAGE_NAME):$(DEV_TAG) AWS/. ; then \
			clear; \
			echo "🛠️  Docker image '$(IMAGE_NAME):$(DEV_TAG)' built successfully!"; \
		else \
			clear; \
			echo "❌ Docker build failed! Check your Dockerfile or network connection."; \
			exit 1; \
		fi

dev-setup:
	@echo "\n⚙️  Initializing DEV environment..."
	docker run \
		--env-file .env \
		-v ./Logs:/app/Logs \
		$(IMAGE_NAME):$(DEV_TAG) \
		./deploy.sh dev

dev-run:
	@echo "\n🚀 Launching DEV container..."
	docker run -it --rm \
		--env-file .env \
		-v ./Logs:/app/Logs \
		$(IMAGE_NAME):$(DEV_TAG) \
		/bin/sh

dev-re: clean dev


# ==========================================
# Production Environment
# ==========================================

prod: prod-build clear-screen connection-test dev-run

prod-build:
	@echo "🛠️  Building Docker image in the PROD environment..."
	@if docker build \
		--progress quiet \
		--target $(DEV_TAG) \
		-t $(IMAGE_NAME):$(DEV_TAG) . >/dev/null 2>&1; then \
			clear; \
			echo "🚀 Docker image '$(IMAGE_NAME):$(DEV_TAG)' built successfully!"; \
		else \
			clear; \
			echo "❌ Docker build failed! Check your Dockerfile or network connection."; \
			exit 1; \
		fi
# prod-run:


# ==========================================
# General
# ==========================================

connection-test:
	@echo "\n🔑 Testing connection to AWS CLI..."
	@OUTPUT=$$(docker run --rm --env-file .env $(IMAGE_NAME):$(DEV_TAG) 2>/dev/null); \
	if [ $$? -eq 0 ] && [ ! -z "$$OUTPUT" ]; then \
		USER_ID=$$(echo "$$OUTPUT" | jq -r '.UserId'); \
		echo "📡 User '$$USER_ID' connected successfully!"; \
	else \
		echo "🔐 Failed to connect to AWS. Check your credentials on the .env file."; \
		exit 1; \
	fi

clean-aws:
	@echo "Cleaning up AWS services..."
	RUNNING_CONTAINER=$$(docker ps --filter "ancestor=$(IMAGE_NAME):$(DEV_TAG)" --format "{{.Names}}"); \
	if [ -z "$$RUNNING_CONTAINER" ]; then \
		docker run --rm --env-file .env $(IMAGE_NAME):$(DEV_TAG) /app/cleanup.sh; \
	else \
		docker exec --env-file .env $$RUNNING_CONTAINER /app/cleanup.sh; \
		docker stop $$RUNNING_CONTAINER; \
	fi
	@echo "AWS Services cleaned up."

clean-docker:
	@echo "🧹 Cleaning up Docker infrastructure..."
	docker system prune -a
	docker system df

clean: clean-aws clean-docker

.PHONY: dev dev-build dev-setup dev-run dev-re \
	prod prod-build prod-run prod-re \
	connection-test clean-aws clean-docker clean
