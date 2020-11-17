PROJECT = septapig
NAME = eventer
TAG = dev

ifndef $(GOPATH)
  export GOPATH=${HOME}/gopath
  ${shell mkdir -p ${GOPATH}}
endif

ifndef $(GOBIN)
  export GOBIN=${GOPATH}/bin
endif


.PHONY: small
small:
	kind delete cluster
	kind create cluster --config calico/kind-calico-small.yaml
	kubectl apply -f calico/ingress-nginx.yaml
	kubectl apply -f calico/tigera-operator.yaml
	kubectl apply -f calico/calicoNetwork.yaml
	kubectl apply -f calico/calicoctl.yaml
	kubectl apply -f calico/cert-manager.yaml



.PHONY: kafka
kafka:
	kind delete cluster
	kind create cluster --config calico/kind-calico-pvc.yaml
	kubectl apply -f calico/ingress-nginx.yaml
	kubectl apply -f calico/tigera-operator.yaml
	kubectl apply -f calico/calicoNetwork.yaml
	kubectl apply -f calico/calicoctl.yaml
	kubectl apply -f calico/cert-manager.yaml

# kubectl kudo init
# kubectl kudo install zookeeper
# kubectl kudo install kafka
# kubectl apply --kustomize github.com/kubernetes/ingress-nginx/deploy/prometheus/
# kubectl apply --kustomize github.com/kubernetes/ingress-nginx/deploy/grafana/


# Or maybe:
# https://github.com/prometheus-operator/prometheus-operator


# kubectl get svc -n ingress-nginx
# Ref: https://kubernetes.github.io/ingress-nginx/user-guide/monitoring/
# kubectl get nodes --selector=kubernetes.io/role!=master -o jsonpath={.items[*].status.addresses[?\(@.type==\"InternalIP\"\)].address}

# 172.18.0.2:30815


