# WebChatBackend

# Setup Project
- Clone repository
- Run `docker build -t golang-api .`
- Run `docker run -p 8080:8000 golang-api`

# Usage
- In the web browser open {host}:8080

# Tests
- Run `make test`

# Project Structure
The project is divided in 3 layers

The Web layer handled by controllers to process web requests.

The Service layer that processes the business logic of the application.

The Repository layer that provides an interface to the application data.