#
MAKEFLAGS += -s
IMAGE_NAME = sturdy-tribble
DEV_TAG = dev
PROD_TAG = prod

ifneq ($(wildcard .env),)
    include .env
    export
endif

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
		$(IMAGE_NAME):$(DEV_TAG) \
		/app/deploy.sh dev

dev-run:
	@echo "\n🚀 Launching DEV container..."
	docker run -it --rm \
		--env-file .env \
		$(IMAGE_NAME):$(DEV_TAG) \
		/bin/sh

# Clean up local docker images
dev-clean:
	@echo "🧹 Cleaning up Docker images..."
	docker system prune -a
	docker system df

dev-re: dev-clean dev

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

# Deletes all the aws service instances and spins them again back up
fclean:

.PHONY: dev dev-build dev-run dev-setup dev-clean dev-re \
	prod prod-build prod-run prod-clean prod-re \
	connection-test fclean
