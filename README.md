# kube-manager

kube-manager is a CLI to import and export Kubernetes contexts to and from a kubeconfig file. 

Ever faced this problem when you downloaded the kubeconfig file for a cluster from a cloud provider (or someone on your team shared theirs with you to give you access) and you want that cluster added to your main kubeconfig file (`~/.kube/config`) where all your other contexts are? You had to manually copy the context from the new kubeconfig file to your existing one. This CLI solves that problem by allowing you to import and export contexts to and from a kubeconfig file.

![kube-manager demo](https://github.com/RinkiyaKeDad/kube-manager/blob/main/demo.gif)

## Installation

### Homebrew
Add the tap:
```sh
brew tap rinkiyakedad/rinkiyakedad
```
Get the package:
```sh
brew install kube-manager
```

### Source Code
1. Build the binary by running `go build` from the root of this repository after cloning it.
2. Move the binary to a directory in your `$PATH` (e.g. `/usr/local/bin`).

## Usage

1. To choose your main kubeconfig (the one from which you want to export contexts/into which you want to import contexts), run `kube-manager kubeconfig <path to your kubeconfig>`. By default the CLI will use the kubeconfig present at: `$HOME/.kube/config`. 

1. To export a context from your main kubeconfig, run `kube-manager export`. This will show you the list of contexts in your main kubeconfig. Select the one you want to export and it will create a new kubeconfig file in the current directory with the selected context.

1. To import a context to your main kubeconfig, run `kube-manager import <path to the kubeconfig file you want to import from>`. This will all all the contexts in that file to your main kubeconfig.

## Contributing

If you like the project, please consider giving it a star. If you have any feedback, please open an issue.

## References

This which helped me create this project.

- https://askcloudarchitech.com/posts/tutorials/create-homebrew-tap-golang-goreleaser-cobra-cli/