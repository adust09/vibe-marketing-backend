FROM mcr.microsoft.com/devcontainers/go:1.23

# Install additional tools
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
        postgresql-client \
        redis-tools \
        curl \
        wget \
        git \
        make \
        build-essential \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install github.com/golang/mock/mockgen@latest

# Set up workspace
WORKDIR /workspace

# Create non-root user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && chown -R $USERNAME:$USERNAME /go

USER $USERNAME