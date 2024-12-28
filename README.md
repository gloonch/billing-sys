# Billing System Project

This project is a billing system designed to manage buildings, units, and payments. The project uses a modular architecture and employs various design patterns and modern software engineering practices to ensure scalability, maintainability, and extensibility.

---

## **Features**

- **CRUD Operations**: Fully functional CRUD operations for Buildings, Units, and Payments.
- **Design Patterns**:
    - **Strategy Pattern**: Used for calculating building charges based on different criteria (e.g., area-based, occupant-based).
    - **Decorator Pattern**: Implemented for adding logging functionality to repositories.
- **Middleware**:
    - **Authentication**: Secure access to the API with JWT-based authentication middleware.
    - **Logger**: Structured and categorized logging for improved debugging.
- **Domain-Driven Design (DDD)**:
    - Modular architecture with separation of domain, application, and infrastructure layers.
- **Caching Layer**: Optimized data retrieval using Redis caching.
- **Event-Driven Architecture**: Asynchronous processing of tasks and events using message queues .
- **Metrics and Monitoring**:
    - **Prometheus**: Collects real-time metrics for the system.
    - **Grafana**: Visualizes collected metrics on custom dashboards.
    - **Tracing**: Distributed tracing with Jaeger to track API calls and dependencies.
- **CI/CD**: Automated build, test, and deployment pipeline using GitHub Actions.

---

## **How to Run the Project**

### **Prerequisites**

1. Install [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
2. Clone the repository:
   ```bash
   git clone https://github.com/your-username/billing-system.git
   cd billing-sys
   docker compose up --build

###

### **Access the services**

API: http://localhost:8000

Prometheus: http://localhost:9090

Grafana: http://localhost:3000 (Default credentials: admin/admin)
