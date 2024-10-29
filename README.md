# Kubernetes Auto Scale Queue Service

## Overview

A scalable system using Kubernetes and Docker. The system is designed to process tasks sent from a simple queueing system. The system can scale automatically based on the load or volume of incoming requests.

## Key features

- Set up a Kubernetes cluster that can manage a Docker container instance.
- Create and deploy a Docker image that processes tasks sent from a simple queueing
  system.
- Ensure the system can scale automatically based on the load or volume of incoming
  requests.

## Getting started

### Prerequisites

It is recommended to have Go installed before running the project. Go can be installed from the official [Go website](https://go.dev/doc/install). This project was built using Go version `1.23`.

### Note

Their is one publisher to push events to RabbitMQ and one consumer to consume events from that queue. The consumer is designed to be scalable and can be scaled based on the load or volume of incoming requests.

### Steps to run

The project dependencies need to be in place for running the project. To install the dependencies, inside the project directory, run:

```go
go mod tidy
```

Once the project dependencies are installed, the project can be run using:

```go
go run main.go
```