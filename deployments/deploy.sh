#!/bin/bash
kubectl version --client
kubectl -n staging get po
kubectl -n staging delete po $(kubectl get pods --namespace staging -l "app=backend" -o jsonpath="{.items[0].metadata.name}")
kubectl apply -f deployments/kubernetes/backend-deployment.yaml
kubectl -n staging get po
