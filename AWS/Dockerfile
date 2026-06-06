# ==========================================
# STAGE 1: Development Environment
# ==========================================

FROM alpine:3.18 AS dev

RUN apk add --no-cache \
    curl \
    unzip \
    groff \
    less \
    python3 \
    py3-pip \
    nodejs \
    npm \
    go \
    nano \
    && pip3 install --no-cache-dir awscli --break-system-packages
    # postgresql
    # redis

WORKDIR /app
COPY deploy.sh .
COPY ./config .

CMD ["aws", "sts", "get-caller-identity"]


# ==========================================
# STAGE 2: Development Environment
# ==========================================

FROM alpine:3.18 AS prod

RUN apk add --no-cache \
    curl \
    unzip \
    groff \
    less \
    python3 \
    py3-pip \
    nodejs \
    npm \
    go \
    && pip3 install --no-cache-dir awscli --break-system-packages

WORKDIR /app

CMD ["aws", "--version"]
