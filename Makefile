#
MAKEFLAGS		+= -s

TIMESTAMP		:= $(shell date +%Y-%m-%d)
LOGFILE 		:= ./Logs/Build/build_$(TIMESTAMP).log
CURRENT_FILE	:= $(lastword $(MAKEFILE_LIST))

MOVE_UP   		:= \033[1A
CLEAR_LN  		:= \033[2K

IMAGE_NAME		= sturdy-tribble
DEV_TAG			= dev
PROD_TAG		= prod

ifneq ($(wildcard .env),)
    include .env
    export
endif


# ==========================================
# Development Environment
# ==========================================

dev: dev-build connection-test dev-setup dev-run

dev-build: init-log
	@echo "[.][🔨] Building Docker image '$(IMAGE_NAME):$(DEV_TAG)'..."
	$(call init_log,build)
	@if docker build \
		--target $(DEV_TAG) \
		--no-cache \
		-t $(IMAGE_NAME):$(DEV_TAG) AWS/. ; then \
			clear; \
			echo "[✅][🔨]  Building Docker image '$(IMAGE_NAME):$(DEV_TAG)'..."; \
		else \
			clear; \
			echo "❌ Docker build failed! Check your Dockerfile or network connection."; \
			exit 1; \
		fi

dev-setup:
	@echo "\n[.][⚙️] Initialize DEV environment..."
	$(call logger,INFO,Initializing DEV environment...,Makefile,0);
	docker run \
		--env-file .env \
		-v ./Logs:/app/Logs \
		$(IMAGE_NAME):$(DEV_TAG) \
		./deploy.sh dev
	@echo "$(MOVE_UP)$(CLEAR_LN)[✅][⚙️] Initialize DEV environment";

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
	@echo "[.][🛠️]  Building Docker image in the PROD environment..."
	@if docker build \
		--progress quiet \
		--target $(DEV_TAG) \
		-t $(IMAGE_NAME):$(DEV_TAG) . >/dev/null 2>&1; then \
			clear; \
			echo "$(MOVE_UP)$(CLEAR_LN)[✅][🛠️] Docker image '$(IMAGE_NAME):$(DEV_TAG)' built successfully!"; \
			$(call logger,INFO,Docker image '$(IMAGE_NAME):$(DEV_TAG)' built successfully!,Makefile,0); \
		else \
			clear; \
			echo "$(MOVE_UP)$(CLEAR_LN)[❌][🛠️] Failed to build Docker container."; \
			$(call logger,INFO,Docker build failed! Check your Dockerfile or network connection.,Makefile,0); \
			exit 1; \
		fi
# prod-run:


# ==========================================
# General
# ==========================================

connection-test:
	@echo "[.][📡] Connection test to AWS CLI."
	$(call logger,INFO,Connection test to AWS CLI...,Makefile,0)
	@OUTPUT=$$(docker run --rm --env-file .env $(IMAGE_NAME):$(DEV_TAG) 2>/dev/null); \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ] && [ ! -z "$$OUTPUT" ]; then \
		USER_ID=$$(echo "$$OUTPUT" | jq -r '.UserId'); \
		printf "$(MOVE_UP)$(CLEAR_LN)[✅][📡] User '$$USER_ID' connected successfully!\n"; \
		$(call logger,INFO,User '$$USER_ID' connected successfully!,Makefile,0); \
	else \
		printf "$(MOVE_UP)$(CLEAR_LN)[❌][📡] Failed to connect to AWS.\n"; \
		$(call logger,INFO,Failed to connect to AWS. Check your credentials on the .env file.,Makefile,0); \
		exit 1; \
	fi

init-log:
	@echo "\n[.][📝] Initializing logging engine..."
	@mkdir -p ./Logs/Build
	@if [ ! -f "$(LOGFILE)" ]; then \
		touch "$(LOGFILE)"; \
		chmod -w "$(LOGFILE)"; \
	fi
	@echo "{'timestamp': '$$(date +%Y-%m-%d_%H:%M:%S)', 'message': 'Start building...', 'process': 'Makefile', 'exit': '0'}" >> $(LOGFILE)
	@echo "$(MOVE_UP)$(CLEAR_LN)[✅][📝] Initializing logging engine..."

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


# ==========================================
# Defines
# ==========================================

# $(1) = LEVEL
# $(2) = MESSAGE
# $(3) = PROCESS
define logger
	echo "{\"timestamp\": \"$$(date +%Y-%m-%d_%H:%M:%S)\", \"level\": \"$(1)\", \"message\": \"$(2)\", \"process\": \"$(3)\"}" >> $(LOGFILE)
endef

.PHONY: dev dev-build dev-setup dev-run dev-re \
	prod prod-build prod-run prod-re \
	connection-test clean-aws clean-docker clean
