# cicd-dev-environment

# Demo Project: Integrating Tilt, kind, Argo Workflows & Events, and ArgoCD

Welcome to this demo project! This repository is designed as a hands-on demonstration of how you can streamline your development workflow. We've integrated Tilt, kind, Argo Workflows & Events, and ArgoCD, providing a seamless solution that allows developers to replicate all steps of the workflow from writing code to deploying it using ArgoCD.

## About the Project

This project focuses on bridging the gap between software development and operations using DevOps' best tools. 
We use `Tilt` to ensure your app is efficiently built and deployed when working locally,
along with `kind` to act as a local Kubernetes for testing.
The heavy-lifting of CI/CD is handled by `Argo Workflows & Events` and `ArgoCD`, giving you fast, reproducible pipelines for your applications.

Let's deep dive into the project, and explore how you can utilize these technologies to supercharge your development workflow.

Stay tuned for more updates!

## Tools in use
### Tilt

**[Tilt](https://tilt.dev/)** is a multi-service development environment tool. It helps you build and deploy your applications on Kubernetes while you're developing locally. You can easily visualize and investigate how your services start up and update in real-time in a web-based dashboard.

### kind

**[kind (Kubernetes in Docker)](https://kind.sigs.k8s.io/)** is a tool for running local Kubernetes clusters using Docker container nodes. It is primarily designed for testing, CI, and development. kind is lightweight, quick, and is part of the Kubernetes project, ensuring up-to-date support.

### ctlptl

**[ctlptl](https://github.com/tilt-dev/ctlptl)** is a command-line tool for declaratively setting up local Kubernetes clusters for development. By integrating with existing tools such as kind, minikube, etc., it provides a smoother and consistent experience for managing local clusters, focusing on configuration in a way that's friendly to scripting and automation.

### Argo Workflows & Events

**[Argo Workflows](https://argoproj.github.io/argo-workflows/)** is an open-source workflow management system that allows you to orchestrate parallel jobs on Kubernetes. It is designed from the ground up for containers, helping you streamline your pipelines for application development and deployment.

**[Argo Events](https://argoproj.github.io/argo-events/)** is an event-based dependency manager for Kubernetes that helps define, create, and manage derivative actions, chains, or dependencies from events occurring in Kubernetes or elsewhere.

### ArgoCD

**[ArgoCD](https://argoproj.github.io/argo-cd/)** is a declarative, GitOps continuous delivery tool for Kubernetes. It allows you to define, deploy, and monitor applications directly from a git repository, ensuring that the state of your live system  matches the desired state in your version control system.

## Getting Started

The easiest way to get started with this project will be to consult its blog post here -> https://anasdomesticus.hashnode.dev/revolutionising-ci-with-argo-events-workflows-tilt-faster-real-time-developer-feedback