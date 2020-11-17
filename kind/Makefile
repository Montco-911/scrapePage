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
