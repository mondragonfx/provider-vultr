# Provider Vultr

`provider-vultr` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Vultr API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/vultr/provider-vultr):
```
up ctp provider install vultr/provider-vultr:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-vultr
spec:
  package: vultr/provider-vultr:v0.1.0
EOF
```
```
git clone git@github.com:mondragonfx/provider-vultr.git
cd provider-vultr
```
Apply the CRDs
```
kubectl apply -f package/crds
```

Run the provider:
```
make run
```

Apply ProviderConfig and example manifests (In another terminal since the previous command is blocking):
```

# Create "crossplane-system" namespace if not exists
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -

#Add your VULTR_API_KEY to examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/
kubectl apply -f examples/kubernetes/kubernetes.yaml
```

Observe managed resources and wait until they are ready:
```
watch kubectl get managed
```

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/vultr/provider-vultr/issues).
