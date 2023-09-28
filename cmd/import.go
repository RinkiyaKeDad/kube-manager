package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func ImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import Kubernetes contexts from file",
		Run: func(cmd *cobra.Command, args []string) {
			importKubeconfig(args)
		},
	}

	return cmd
}

func importKubeconfig(args []string) {
	if len(args) != 1 {
		fmt.Println("Please provide the path to the kubeconfig file to import")
		os.Exit(1)
	}

	importedPath := args[0]
	importedConfig, err := clientcmd.LoadFromFile(importedPath)
	if err != nil {
		fmt.Println("Failed to load kubeconfig:", err)
		os.Exit(1)
	}

	config, err := loadKubeconfig()
	if err != nil {
		fmt.Println("Failed to load kubeconfig:", err)
		os.Exit(1)
	}

	for context, contextConfig := range importedConfig.Contexts {
		config.Contexts[context] = contextConfig
	}

	for cluster, clusterConfig := range importedConfig.Clusters {
		config.Clusters[cluster] = clusterConfig
	}

	for authInfo, authInfoConfig := range importedConfig.AuthInfos {
		config.AuthInfos[authInfo] = authInfoConfig
	}

	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		fmt.Println("Failed to import kubeconfig:", err)
		os.Exit(1)
	}

	fmt.Println("Kubeconfig imported successfully")
}
