# gin-skeleton

## Clean Architecture Overview

This project follows the principles of Clean Architecture, which emphasizes separation of concerns and independence of frameworks, databases, and external systems. The architecture is divided into layers, each with specific responsibilities:

### 1. **Domain Layer**

- **Purpose**: Represents the core business logic and entities.
- **Location**: `internal/domain`
- **Details**:
  - Contains the `User` struct, which defines the core attributes of a user.
  - This layer is independent of any other layer and does not depend on external libraries.

### 2. **Use Case Layer**

- **Purpose**: Contains the application-specific business rules.
- **Location**: `internal/usecase`
- **Details**:
  - Implements the logic for creating and retrieving users.
  - Interacts with the repository layer to fetch or persist data.
  - Example: `UserUsecase` handles user-related operations.

### 3. **Repository Layer**

- **Purpose**: Handles data persistence and retrieval.
- **Location**: `internal/repository`
- **Details**:
  - Uses GORM to interact with the database.
  - Example: `UserRepository` provides methods to create and fetch users from the database.

### 4. **Delivery Layer**

- **Purpose**: Handles the input/output of the application.
- **Location**: `internal/delivery`
- **Details**:
  - **HTTP Handlers**: Located in `internal/delivery/http`, responsible for handling HTTP requests and responses.
  - **Middleware**: Located in `internal/delivery/middleware`, responsible for cross-cutting concerns like error handling, authentication, and CORS.

### 5. **Configuration Layer**

- **Purpose**: Manages application configuration and setup.
- **Location**: `config`
- **Details**:
  - Handles environment variables, database connections, and logger setup.

### 6. **Command Layer**

- **Purpose**: Initializes dependencies and sets up routes.
- **Location**: `cmd`
- **Details**:
  - `cmd/routes`: Defines API routes.
  - `cmd/Handler`: Configures handlers with their dependencies.

### Key Principles

- **Dependency Rule**: Inner layers should not depend on outer layers. For example, the domain layer does not depend on the repository or delivery layers.
- **Testability**: Each layer can be tested independently, ensuring high test coverage.
- **Framework Independence**: The architecture is not tied to any specific framework, making it easier to replace frameworks if needed.

### Project Structure

```
gin/
├── cmd/                # Command layer for initializing dependencies and routes
├── config/             # Configuration layer for environment variables and database setup
├── internal/
│   ├── delivery/       # Delivery layer for HTTP handlers and middleware
│   ├── domain/         # Domain layer for core business entities
│   ├── repository/     # Repository layer for data persistence
│   ├── usecase/        # Use case layer for business logic
├── pkg/                # Shared utilities (e.g., logger, JWT, validators)
├── test/               # Unit and integration tests
└── main.go             # Entry point of the application
```

This structure ensures a clean separation of concerns, making the codebase maintainable, scalable, and easy to understand.
