# GoConnect

A scalable social media backend service built with Go, implementing modern architectural patterns and industry best practices for high-performance distributed systems.

## Overview

GoConnect is a production-ready social media backend that demonstrates enterprise-level software engineering practices. The system is designed with Clean Architecture principles, employs event-driven patterns, and leverages containerization and orchestration technologies to ensure scalability and maintainability.

## Architecture

The application follows **Clean Architecture** and **Hexagonal Design** patterns, ensuring clear separation of concerns and testability:

- **Domain Layer**: Core business logic and entities
- **Application Layer**: Use cases and application services  
- **Infrastructure Layer**: External dependencies and frameworks
- **Interface Layer**: API endpoints and external integrations

## Technology Stack

### Core Technologies
- **Runtime**: Go 1.24+
- **Database**: PostgreSQL (primary data store)
- **Cache**: Valkey (Redis-compatible caching layer)
- **API Framework**: Chi router with net/http

### Infrastructure & DevOps
- **Containerization**: Docker & Docker Compose
- **Orchestration**: Kubernetes
- **CI/CD**: GitHub Actions
- **Migration Management**: golang-migrate

### Planned Integrations
- **Object Storage**: Amazon S3
- **Monitoring**: Prometheus & Grafana
- **Tracing**: OpenTelemetry

## Features

### Core Functionality
- **Authentication & Authorization**: JWT-based authentication with role-based access control
- **User Management**: Registration, profile management, and user relationships
- **Content Management**: Post creation, editing, and deletion
- **Social Features**: Like/unlike posts, follow/unfollow users
- **Real-time Features**: Event-driven architecture with pub/sub messaging

### Performance & Scalability
- **Caching Strategy**: Redis-based caching for frequently accessed data
- **Database Optimization**: Efficient PostgreSQL schema design with proper indexing
- **Connection Pooling**: Optimized database and cache connections
- **Health Checks**: Comprehensive service health monitoring

## Quick Start

### Prerequisites
- Go 1.24 or higher
- Docker and Docker Compose
- kubectl (for Kubernetes deployment)

### Local Development
```bash
# Clone the repository
git clone https://github.com/yourusername/goconnect.git
cd goconnect

# Start all services
docker compose up --build

# The API will be available at http://localhost:8080
```

### Kubernetes Deployment
```bash
# Apply Kubernetes manifests
kubectl apply -f k8s/

# Verify deployment
kubectl get pods
kubectl get services
```

## API Documentation

### Authentication
```bash
# Register a new user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"user","email":"user@example.com","password":"password"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'
```

### Posts
```bash
# Create a post (requires authentication)
curl -X POST http://localhost:8080/api/posts \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello, GoConnect!"}'
```

## Development

### Database Migrations
```bash
# Run migrations
migrate -path ./migrations -database "postgres://user:password@localhost/goconnect?sslmode=disable" up

# Create new migration
migrate create -ext sql -dir ./migrations -seq add_new_table
```

### Testing
```bash
# Run unit tests
go test ./...

# Run integration tests
go test -tags=integration ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## CI/CD Pipeline

The project includes a comprehensive CI/CD pipeline using GitHub Actions:

1. **Build Stage**: Compiles Go binary and runs tests
2. **Docker Stage**: Builds and tags container images
3. **Kubernetes Stage**: Deploys to Minikube cluster
4. **Validation Stage**: Runs health checks and integration tests
5. **Migration Stage**: Applies database schema changes

## Monitoring and Observability

### Health Checks
- Database connectivity monitoring
- Cache availability verification
- Service dependency health validation

### Metrics (Planned)
- Request latency and throughput
- Database query performance
- Cache hit/miss ratios
- Error rates and patterns

## Security

- JWT token-based authentication
- Password hashing with bcrypt
- Input validation and sanitization
- CORS configuration
- Rate limiting (planned)

## Performance Considerations

- Connection pooling for database and cache
- Efficient query patterns with proper indexing
- Caching strategies for frequently accessed data
- Graceful degradation under load

## Deployment

### Local Development
Uses Docker Compose for easy local development with all dependencies.

### Production (Kubernetes)
- Horizontal pod autoscaling
- Service discovery and load balancing
- Persistent volume management
- Health probes and readiness checks

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [golang-migrate](https://github.com/golang-migrate/migrate) for database migration management
- [Chi](https://github.com/go-chi/chi) for HTTP routing
- [Valkey](https://valkey.io/) for Redis-compatible caching

---

**Author**: Suresh S  
**Contact**: [Your Email]  
**LinkedIn**: [Your LinkedIn Profile]

> Built with modern Go practices and enterprise-grade architecture patterns.