<h1 align="center">Go Bank 🏦</h1>

This application is primarily intended for exploring technical concepts. My goal is to experiment with different technologies, software architecture designs, and all the essential components involved in building a monolithic-based application in Golang.

## Features :sparkles:

- ✅ Building RESTful APIs using the `Gin` framework for handling HTTP requests
- ✅ Leveraging `gRPC` for efficient internal service-to-service communication
- ✅ Secure authentication using `PASETO` tokens and `JWT`
- ✅ Asynchronous task processing with `Redis` and `Asynq` for background processing
- ✅ Reliable database operations with `PostgreSQL` and `pgx` driver
- ✅ Type-safe database queries with `SQLC`
- ✅ Database migrations using `golang-migrate`
- ✅ Input validation using `go-playground/validator`
- ✅ Structured logging with `zerolog`
- ✅ Configuration management with `Viper`
- ✅ Testing with `testify` for assertions and `gomock` for mocking
- ✅ Containerized deployment with `Docker` and `docker-compose`
- ✅ Deploying to production environments using `Kubernetes` on `AWS EKS`
- ✅ Automatic SSL/TLS certificate management with `cert-manager` and `Let's Encrypt`
- ✅ CI/CD pipeline with `GitHub Actions`

## Technologies - Libraries 🛠️

<h4>Core Framework</h4>

- **[gin-gonic/gin](https://github.com/gin-gonic/gin)** - High-performance HTTP web framework
- **[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)** - gRPC to JSON proxy

<h4>Database & Migration</h4>

- **[jackc/pgx](https://github.com/jackc/pgx)** - PostgreSQL driver and toolkit
- **[golang-migrate/migrate](https://github.com/golang-migrate/migrate)** - Database migrations
- **[sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)** - Type-safe SQL query builder

<h4>Validation</h4>

- **[go-playground/validator](https://github.com/go-playground/validator)** - Go Struct and Field validation

<h4>Authentication & Security</h4>

- **[o1egl/paseto](https://github.com/o1egl/paseto)** - Platform-Agnostic Security Tokens
- **[dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)** - JSON Web Tokens
- **[golang/crypto](https://github.com/golang/crypto)** - Cryptographic functions

<h4>Task Processing</h4>

- **[hibiken/asynq](https://github.com/hibiken/asynq)** - Distributed task queue
- **[redis/go-redis](https://github.com/redis/go-redis)** - Redis client

<h4>Logging</h4>

- **[rs/zerolog](https://github.com/rs/zerolog)** - Zero-allocation JSON logger

<h4>Testing & Mocking</h4>

- **[stretchr/testify](https://github.com/stretchr/testify)** - Testing toolkit
- **[uber-go/mock](https://github.com/uber-go/mock)** - Mocking framework

<h4>Documentation</h4>

- **[Swagger](https://swagger.io/)** - API documentation with static files for Swagger UI
- **[DBML](https://www.dbml.org/)** - Database Markup Language for visualizing database schema

<h4>Configuration & Environment</h4>

- **[spf13/viper](https://github.com/spf13/viper)** - Configuration solution

<h4>Deployment & DevOps</h4>

- **[AWS EKS](https://aws.amazon.com/eks/)** - Managed Kubernetes service
- **[Kubernetes](https://kubernetes.io/)** - Container orchestration
- **[Docker](https://www.docker.com/)** - Container platform
- **[Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/)** - Kubernetes ingress controller
- **[cert-manager](https://cert-manager.io/)** - Certificate management for Kubernetes
- **[Let's Encrypt](https://letsencrypt.org/)** - Free SSL/TLS certificates

## System Architecture 🏗️

The application follows a monolithic architecture pattern with some clean architecture principles applied.
