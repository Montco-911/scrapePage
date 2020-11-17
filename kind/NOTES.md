

```
VERSION=0.17.2
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
wget -O kubectl-kudo https://github.com/kudobuilder/kudo/releases/download/v${VERSION}/kubectl-kudo_${VERSION}_${OS}_${ARCH}
chmod +x kubectl-kudo
# add to your path
sudo mv kubectl-kudo /usr/local/bin/kubectl-kudo

```