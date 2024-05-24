# echo-clean-arc

This repository demonstrates how to progressively build a REST API application using the Echo framework in Go. We start with a very simple Echo app and gradually refactor it into a clean architecture, adding more complex structures and external components like Redis, Kafka, RabbitMQ, gRPC clients, and more.

## Project Goal

The main goal of this project is to provide a clear and practical example of how to evolve a simple REST API into a more robust and maintainable application using clean architecture principles and integrating various external services.

## Steps Taken

The project follows these steps:

1. **Simple Echo App:**  Begin with a basic Echo application with minimal structure.
2. **Database Integration:** Integrate a database (SQLite in this example) using Gorm as the ORM.
3. **Refactoring to Clean Architecture:** Introduce the concepts of Clean Architecture, separating concerns into distinct layers (domain, repository, service, handler).
4. **Dependency Injection:** Implement dependency injection using Wire to improve modularity and testability.
5. **External Services:**  Gradually add integrations with external services like:
   - PostgreSQL for the database
   - Redis for caching
   - gRPC clients for inter-service communication
   - GraphQ
   - Kafka
   - RabbitMQ

## Future Enhancements

- **More External Services:** Integrate with more external services like Redis, Kafka, and RabbitMQ.
- **gRPC Communication:** Implement gRPC clients for communication with other services.
- **Testing:** Add comprehensive unit and integration tests.
- **Documentation:**  Improve documentation and add more explanations.
- **CI/CD:**  Set up a CI/CD pipeline for automated building and deployment.
