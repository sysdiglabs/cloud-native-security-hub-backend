#!/bin/bash
kubectl version --client
kubectl -n promhub get po
kubectl -n promhub delete po -l app=backend
kubectl apply -f deployments/master/backend-deployment.yaml
kubectl -n promhub get po