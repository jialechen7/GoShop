# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Build and Development
- `make build` - Build the application using gf CLI with configuration from hack/config.yaml
- `make up` - Update GoFrame and its CLI to latest stable version
- `gf run` - Run the application in development mode (after installing gf CLI)

### Code Generation
- `make dao` - Generate Go files for DAO/DO/Entity from database schema
- `make ctrl` - Parse API and generate controller/SDK files
- `make service` - Generate Go files for Service layer
- `make enums` - Parse project files and generate enums

### Database and Infrastructure
- Database connection: MySQL at 127.0.0.1:3306/goshop (see manifest/config/config.yaml)
- Redis connection: 127.0.0.1:6379 db:1 (see manifest/config/config.yaml)
- Create Redis streams for async processing: `XGROUP CREATE stream.user_coupon g1 0 MKSTREAM`

### Docker and Deployment
- `make image` - Build docker image with git-based tagging
- `make image.push` - Build and push docker image
- `make deploy` - Deploy using Kustomize to current kubectl environment

## Architecture Overview

### Project Structure
This is a GoFrame-based e-commerce system with clear separation between frontend and backend APIs:

- **API Layer**: Separate packages for `backend` (admin) and `frontend` (customer) APIs in `/api`
- **Controller Layer**: HTTP handlers in `/internal/controller`
- **Logic Layer**: Business logic in `/internal/logic` with domain-specific packages
- **Service Layer**: Service interfaces and implementations in `/internal/service`
- **DAO Layer**: Database access objects in `/internal/dao` with auto-generated code
- **Model Layer**: Data models with `entity` (DB structs), `do` (query objects), and domain models

### Key Features
- **Dual Authentication**: Separate gtoken authentication for admin backend and user frontend
- **Redis Integration**: Used for caching, distributed locking (redsync), and streaming (user coupon processing)
- **Async Processing**: Redis streams for handling user coupon operations asynchronously
- **Seckill System**: Flash sale functionality with Lua scripts for atomic operations
- **File Upload**: Support for local and Qiniu cloud storage
- **OAuth Integration**: GitHub OAuth support for admin login

### Database Integration
- Uses GoFrame's ORM with auto-generated DAO/DO/Entity files
- Database configuration supports connection pooling and debugging
- Soft delete functionality with automatic timestamp management

### Key Configuration Files
- `manifest/config/config.yaml` - Main application configuration
- `hack/config.yaml` - CLI tool configuration for code generation
- `go.mod` - Dependencies including GoFrame v2.7.4, Redis clients, JWT, and utilities

### Development Workflow
1. Database schema changes: Update database, then run `make dao` to regenerate models
2. API changes: Update API definitions, then run `make ctrl` to regenerate controllers
3. Service layer changes: Run `make service` to regenerate service interfaces
4. Build and test: Use `make build` for production builds

### Security Notes
- Uses gtoken for JWT-based authentication with separate admin/user contexts
- CORS middleware configured for cross-origin requests
- Input validation using GoFrame's validation tags
- Distributed Redis locks for preventing race conditions in seckill operations