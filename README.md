# Turkish Wordle Game 🎮

A modern implementation of the popular Wordle game, localized for Turkish language speakers. This project is built with
a focus on clean architecture, type safety, and modern web development practices. The backend follows Domain-Driven
Design (DDD) principles for better maintainability and scalability.

## 🚧 Project Status

This project is currently under active development. We're continuously adding new features, improving the user
experience, and optimizing performance.

## 🛠 Tech Stack

### Frontend

- React with TypeScript
- Vite for build tooling
- TailwindCSS for styling
- ESLint + Prettier for code quality

### Backend

- Go with Echo framework
- Domain-Driven Design (DDD) architecture
- PostgreSQL for data storage
- Docker for containerization
- GitHub Actions for CI/CD

### Infrastructure

- AWS ECS for container orchestration
- AWS Lambda for game rotation
- AWS RDS for PostgreSQL
- AWS CloudWatch for monitoring

## 🏗 Architecture

The backend follows Domain-Driven Design principles:

- **Domain Layer**: Core business logic and rules
- **Application Layer**: Use cases and orchestration
- **Infrastructure Layer**: External services, databases, and frameworks
- **Interface Layer**: API endpoints and controllers

## 🎯 Features

- [x] Turkish word support
- [x] Responsive design
- [x] Keyboard input support
- [x] Game state persistence
- [x] Daily word rotation (AWS Lambda)
- [ ] User statistics
- [ ] Social sharing
- [ ] Dark/Light theme
- [ ] Offline support

## 🏗 Project Structure

```
.
├── client/                # Frontend React application
├── server/                # Backend Go/Echo application
│   ├── cmd/              # Application entry points
│   │   ├── api/         # Main API server
│   │   └── lambda/      # AWS Lambda functions
│   ├── internal/        # Private application code
│   │   ├── domain/      # Domain layer (entities, value objects)
│   │   ├── application/ # Application layer (use cases)
│   │   ├── infrastructure/ # Infrastructure layer (repositories, services)
│   │   └── presentation/  # Presentation layer (handlers, routes)
├── docker/               # Docker configuration files
│   ├── game/             # Game dockerfile for lambda
├── ngnix/                 # Nginx configuration files  
```

## 🚀 Getting Started

Detailed setup instructions will be added as the project matures.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the Apache 2.0 License - see the LICENSE file for details.

## ⚠️ Note

This is a work in progress. Features, documentation, and setup instructions will be updated as the project evolves. 