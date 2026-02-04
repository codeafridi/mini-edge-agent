# Mini Edge Agent

A minimal Go-based Kubernetes controller simulator that mimics basic edge agent behavior.

## Overview

This project demonstrates how a lightweight edge-like controller can:
- Watch Kubernetes resources
- React to cluster state changes
- Enable or disable workloads based on connectivity conditions

Inspired by Kubernetes controller patterns and edge computing concepts (e.g., KubeEdge).

## Features

- Written in Go
- Uses `client-go` informers
- Watches Deployments
- Simulates edge online/offline behavior
- Scales workloads automatically

## How It Works

1. The agent connects to the Kubernetes API.
2. It watches Deployment resources using informers.
3. Based on a flag (`--edge-online`), it:
   - Scales deployments to zero when the edge is offline.
   - Leaves them running when online.

## Usage

### Start a local cluster
```bash
minikube start

 Create a test deployment

kubectl create deployment demo --image=nginx

Run the Mini Edge Agent
go run main.go --edge-online=false

Verify behavior
kubectl get deployments


You should see the deployment scaled to 0 replicas when the edge is offline.

Example Scenario

This simulates a situation where an edge node becomes unavailable.
The agent reacts by disabling workloads to prevent inconsistent or unnecessary execution while the edge is offline.

Project Structure
mini-edge-agent/
├── main.go
├── go.mod
├── go.sum
├── config/
│   └── config.yaml

Motivation

This project was created to understand Kubernetes controller mechanics and explore how edge agent behavior can be simulated using core Kubernetes primitives without building a full edge platform.

Future Improvements

Restore original replica counts when the edge comes back online

Filter deployments using labels

Read configuration from a YAML file

Introduce a CRD-based control mechanism

Simulate node-level edge conditions
