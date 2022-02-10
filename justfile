default:
  @just --list

create cluster:
  kind create cluster --config kind.yaml --name demo
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

up:
  skaffold run

down:
  skaffold delete

