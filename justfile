default:
  @just --list

cluster_create:
  kind create cluster --config kind/kind.yaml --name demo
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

cluster_delete:
  kind delete cluster --name demo

up:
  skaffold run

down:
  skaffold delete


