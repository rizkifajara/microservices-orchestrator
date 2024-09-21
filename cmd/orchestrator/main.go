package main

import (
	"fmt"
	"log"

	"microservices-orchestrator/pkg/k8s"
	"microservices-orchestrator/pkg/orchestrator"
)

func main() {
	// Create a new Kubernetes client
	client, err := k8s.NewClient()
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	// Create a new Orchestrator instance
	orch := orchestrator.New(client)

	// Deploy a microservice
	err = orch.DeployMicroservice("my-microservice", "nginx:latest", 3)
	if err != nil {
		log.Fatalf("Failed to deploy microservice: %v", err)
	}

	fmt.Println("Microservice deployed successfully")
}
