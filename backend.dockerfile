# Use Debian 12 as the base image
FROM debian:12

# Set the working directory
WORKDIR /app

# Install necessary dependencies
RUN touch /etc/apt/sources.list && \
    rm -rf /etc/apt/sources.list.d && \
    echo "deb http://ftp.us.debian.org/debian/ bookworm main" > /etc/apt/sources.list && \
    echo "deb http://ftp.us.debian.org/debian/ bookworm-updates main" >> /etc/apt/sources.list && \
    apt-get update && apt-get install -y \
    wget \
    git \
    && rm -rf /var/lib/apt/lists/*

# Install Go 1.22.4

RUN wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz && \
    rm go1.22.4.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the source code
COPY cmd ./cmd
COPY internal ./internal

# Build the Go application
RUN go build -o server ./cmd/api

# Set ARG variables
ARG BREWNIQUE_PORT=8080
ARG BREWNIQUE_ENV=dev
ARG BREWNIQUE_DATABASE_DRIVER=postgres
ARG BREWNIQUE_DATABASE_DSN=postgres://brewnique:localdevpass@localhost/brewnique
ARG BREWNIQUE_DATABASE_MAX_OPEN_CONNS=25
ARG BREWNIQUE_DATABASE_MAX_IDLE_CONNS=25
ARG BREWNIQUE_DATABASE_CONN_MAX_LIFETIME=10m

# Expose the port on which the server will listen
EXPOSE $BREWNIQUE_PORT

# Run the server executable
CMD ["./server"]