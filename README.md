<h1 align="center">Go Bank 🏦</h1>

This application is primarily intended for exploring technical concepts. My goal is to experiment with different technologies, software architecture designs, and all the essential components involved in building a monolithic-based application.

## Features :sparkles:

- ✅ RESTful API built with `Gin` framework for handling HTTP requests
- ✅ gRPC API for efficient internal service communication
- ✅ Secure authentication using `PASETO` tokens and `JWT`
- ✅ Asynchronous task processing with `Redis` and `Asynq`
- ✅ Reliable database operations with `PostgreSQL` and `pgx`
- ✅ Type-safe database queries with `SQLC`
- ✅ Database migrations using `golang-migrate`
- ✅ Input validation using `validator`
- ✅ Structured logging with `zerolog`
- ✅ Configuration management with `Viper`
- ✅ Testing with `testify` and `gomock`
- ✅ Containerized deployment with `Docker` and `docker-compose`
- ✅ `Kubernetes` deployment with `AWS EKS`
- ✅ CI/CD pipeline with `GitHub Actions`

## Technologies - Libraries 🛠️

### Core Framework

- **[Gin](https://github.com/gin-gonic/gin)** - High-performance HTTP web framework
- **[gRPC](https://grpc.io/)** - High-performance RPC framework
- **[Viper](https://github.com/spf13/viper)** - Configuration solution
- **[Zerolog](https://github.com/rs/zerolog)** - Zero-allocation JSON logger

### Database & Migration

- **[PostgreSQL](https://www.postgresql.org/)** - Advanced open-source database
- **[pgx](https://github.com/jackc/pgx)** - PostgreSQL driver and toolkit
- **[SQLC](https://sqlc.dev/)** - Type-safe SQL query builder
- **[golang-migrate](https://github.com/golang-migrate/migrate)** - Database migrations

### Authentication & Security

- **[PASETO](https://github.com/o1egl/paseto)** - Platform-Agnostic Security Tokens
- **[JWT](https://github.com/dgrijalva/jwt-go)** - JSON Web Tokens
- **[Crypto](https://golang.org/x/crypto)** - Cryptographic functions

### Task Processing

- **[Asynq](https://github.com/hibiken/asynq)** - Distributed task queue
- **[Redis](https://redis.io/)** - In-memory data structure store

### Testing

- **[Testify](https://github.com/stretchr/testify)** - Testing toolkit
- **[Gomock](https://github.com/uber-go/mock)** - Mocking framework

### Documentations

- **[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)** - gRPC to JSON proxy
- **[Swagger](https://swagger.io/)** - API documentation
- **[DBML](https://www.dbml.org/)** - Database Markup Language for visualizing database schema

### Cloud & Infrastructure

- **[AWS EKS](https://aws.amazon.com/eks/)** - Managed Kubernetes service
- **[Kubernetes](https://kubernetes.io/)** - Container orchestration
- **[Docker](https://www.docker.com/)** - Container platform
- **[Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/)** - Kubernetes ingress controller

## System Architecture 🏗️

The application follows a monolithic architecture pattern with with some clean architecture principles applied.
