# Kernel Module Management

The Kernel Module Management Operator manages out of tree kernel modules in Kubernetes.

> **Note**: This is a downstream fork of the upstream [kubernetes-sigs/kernel-module-management](https://github.com/kubernetes-sigs/kernel-module-management) repository. This fork contains AMD ROCm-specific modifications and enhancements.

## Getting started

For installation and usage instructions, please refer to the [AMD GPU Operator for AMD Instinct documentation](https://rocm.docs.amd.com/projects/gpu-operator/en/latest/).

## Building Images

This project uses Make to build container images. Below are the available build targets:

### Building the Operator Images

Build the main KMM operator image:

```shell
make docker-build
```

This creates the image `docker.io/rocm/kernel-module-management-operator:dev`

Build the hub manager image:

```shell
make docker-build-hub
```

This creates the image `docker.io/rocm/kernel-module-management-operator-hub:dev`

### Building Component Images

Build the signer image:

```shell
make signimage-build
```

This creates the image `docker.io/rocm/kernel-module-management-signimage:dev`

Build the webhook server image:

```shell
make webhookimage-build
```

This creates the image `docker.io/rocm/kernel-module-management-webhook-server:dev`

Build the worker image:

```shell
make workerimage-build
```

This creates the image `docker.io/rocm/kernel-module-management-worker:dev`

### Building All Images

To build all images at once:

```shell
make docker-build docker-build-hub signimage-build webhookimage-build workerimage-build
```

### Customizing Image Tags

You can customize the image tag and registry by setting environment variables:

```shell
# Set custom image tag
export IMAGE_TAG=v1.0.0
make docker-build

# Set custom registry
export IMAGE_TAG_BASE=my-registry.com/kmm-operator
make docker-build
```

### Saving Images

You can save built images to tar.gz files for offline distribution:

```shell
make docker-save           # Saves operator image
make signimage-save        # Saves signer image
make webhookimage-save     # Saves webhook server image
make workerimage-save      # Saves worker image
```

### Building OLM Bundle

Build the Operator Lifecycle Manager (OLM) bundle:

```shell
make bundle-build
```
