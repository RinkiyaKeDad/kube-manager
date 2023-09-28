package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func ExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export selected kubeconfig context",
		Run: func(cmd *cobra.Command, args []string) {
			exportKubeconfig()
		},
	}

	return cmd
}

func exportKubeconfig() {
	config, err := loadKubeconfig()
	if err != nil {
		fmt.Println("Failed to load kubeconfig:", err)
		os.Exit(1)
	}

	fmt.Println("Available contexts:")
	for context := range config.Contexts {
		fmt.Println(context)
	}

	selectedContext := promptForContext(config.Contexts)

	exportedConfig := api.Config{
		CurrentContext: selectedContext,
		Clusters:       map[string]*api.Cluster{selectedContext: config.Clusters[selectedContext]},
		Contexts:       map[string]*api.Context{selectedContext: config.Contexts[selectedContext]},
		AuthInfos:      map[string]*api.AuthInfo{selectedContext: config.AuthInfos[selectedContext]},
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		os.Exit(1)
	}

	exportedPath := filepath.Join(currentDir, fmt.Sprintf("%s.yaml", selectedContext))
	err = clientcmd.WriteToFile(exportedConfig, exportedPath)
	if err != nil {
		fmt.Println("Failed to export kubeconfig:", err)
		os.Exit(1)
	}

	fmt.Println("Kubeconfig exported successfully to:", exportedPath)
}

func promptForContext(contexts map[string]*api.Context) string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Select the context to export: ")
		scanner.Scan()
		context := scanner.Text()

		if _, ok := contexts[context]; ok {
			return context
		}

		fmt.Println("Invalid context. Please try again.")
	}
}
