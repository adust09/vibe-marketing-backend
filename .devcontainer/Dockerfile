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

# Install Go tools with compatible versions
RUN go install github.com/cosmtrek/air@v1.49.0 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.2 && \
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0 && \
    go install github.com/golang/mock/mockgen@v1.6.0

# Set up workspace
WORKDIR /workspace

# Fix permissions for Go directory
RUN chown -R vscode:vscode /go

USER vscode