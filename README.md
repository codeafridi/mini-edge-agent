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
