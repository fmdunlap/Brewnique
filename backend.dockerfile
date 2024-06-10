# Use Debian 12 as the base image
FROM debian:12

# Set the working directory
WORKDIR /app

# Install necessary dependencies
RUN sed -i 's|http://deb.debian.org/debian/|http://ftp.us.debian.org/debian/|g' /etc/apt/sources.list.d/debian.sources && \
    apt-get update && apt-get install -y \
    golang \
    git \
    && rm -rf /var/lib/apt/lists/*

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the source code
COPY cmd ./cmd
COPY internal ./internal

# Build the Go application
RUN go build -o server ./cmd

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

# Set environment variables
ENV BREWNIQUE_PORT=$BREWNIQUE_PORT \
    BREWNIQUE_ENV=$BREWNIQUE_ENV \
    BREWNIQUE_DATABASE_DRIVER=$BREWNIQUE_DATABASE_DRIVER \
    BREWNIQUE_DATABASE_DSN=$BREWNIQUE_DATABASE_DSN \
    BREWNIQUE_DATABASE_MAX_OPEN_CONNS=$BREWNIQUE_DATABASE_MAX_OPEN_CONNS \
    BREWNIQUE_DATABASE_MAX_IDLE_CONNS=$BREWNIQUE_DATABASE_MAX_IDLE_CONNS \
    BREWNIQUE_DATABASE_CONN_MAX_LIFETIME=$BREWNIQUE_DATABASE_CONN_MAX_LIFETIME

# Run the server executable
CMD ["./server"]