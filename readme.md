# Microservices Orchestrator

This project is a Kubernetes-based microservices orchestrator written in Go. It provides a simple way to deploy and manage microservices in a Kubernetes cluster.

## Features

- Deploy microservices to Kubernetes
- Manage deployments and services
- Interact with Kubernetes API

## Prerequisites

- Go 1.16 or later
- Docker
- Kubernetes cluster (local or remote)
- kubectl configured to interact with your cluster

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/microservices-orchestrator.git
   cd microservices-orchestrator
   ```

2. Build the project:
   ```
   go build -o orchestrator cmd/orchestrator/main.go
   ```

3. Build the Docker image:
   ```
   docker build -t orchestrator:v1 .
   ```

## Usage

1. Apply the necessary RBAC rules:
   ```
   kubectl apply -f rbac.yaml
   ```

2. Deploy the orchestrator:
   ```
   kubectl apply -f orchestrator-deployment.yaml
   ```

3. Check the logs:
   ```
   kubectl logs deployment/orchestrator
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.