# Get Go image from DockerHub.
FROM golang:1.22-alpine AS api

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod .
COPY go.sum .

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-go ./cmd/main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=api /compiler/docker-go .

# Run application and expose port 8080.
EXPOSE 8080
CMD ["./docker-go"]