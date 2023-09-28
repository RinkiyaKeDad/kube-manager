# kube-manager

kube-manager is a CLI to export and import Kubernetes contexts from and to a kubeconfig file. 

It is meant to make it easier to share access to K8s clusters by allowing you to `export` one of your contexts to a separate kubeconfig file which you can send to anyone you want to share access with. They can then `import` the context into their original kubeconfig file and have access to the cluster you shared with them.

## Installation

1. Build the binary by running `go build` from the root of this repository after cloning it.
2. Move the binary to a directory in your `$PATH` (e.g. `/usr/local/bin`).

## Usage

1. To choose your main kubeconfig (the one from which you want to export contexts/into which you want to import contexts), run `kube-manager kubeconfig <path to your kubeconfig>`. By default the CLI will use the kubeconfig present at: `$HOME/.kube/config`. 

1. To export a context from your main kubeconfig, run `kube-manager export`. This will show you the list of contexts in your main kubeconfig. Select the one you want to export and it will create a new kubeconfig file in the current directory with the selected context.

1. To import a context to your main kubeconfig, run `kube-manager import <path to the kubeconfig file you want to import from>`. This will all all the contexts in that file to your main kubeconfig.

## Contributing

If you like the project, please consider giving it a star. If you have any feedback, please open an issue.