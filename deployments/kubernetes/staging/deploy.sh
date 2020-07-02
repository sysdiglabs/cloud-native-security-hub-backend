#!/bin/bash
kubectl version --client
kubectl -n staging get po
kubectl -n staging delete po -l app=backend
kubectl apply -f deployments/kubernetes/staging/backend-deployment.yaml
kubectl -n staging get po
