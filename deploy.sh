#!/bin/bash
echo "Updating local cluster..."
kubectl apply -f ./k8s/deployment.yaml
kubectl rollout restart deployment/prr-app-deployment
echo "Deployment successfully completed!"