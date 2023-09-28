package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var kubeconfigPath string

func KubeconfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubeconfig",
		Short: "Set the location of the kubeconfig file",
		Run: func(cmd *cobra.Command, args []string) {
			setKubeconfigPath(args)
		},
	}

	return cmd
}

func setKubeconfigPath(args []string) {
	if len(args) != 1 {
		fmt.Println("Please provide the path to the kubeconfig file")
		os.Exit(1)
	}

	kubeconfigPath = args[0]
	fmt.Println("Kubeconfig path set to:", kubeconfigPath)
}

func loadKubeconfig() (*clientcmdapi.Config, error) {
	if kubeconfigPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		kubeconfigPath = clientcmd.RecommendedHomeFile
		if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
			kubeconfigPath = filepath.Join(homeDir, ".kube", "config")
		}
	}

	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return config, nil
}
