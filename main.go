package main

import (
	"fmt"
	"os"

	"github.com/RinkiyaKeDad/kube-manager/cmd"
	"github.com/spf13/cobra"
)

var kubeconfigPath string

func main() {
	rootCmd := &cobra.Command{
		Use:   "kube-manager",
		Short: "Manage kubeconfig",
	}

	rootCmd.AddCommand(cmd.ExportCmd(), cmd.ImportCmd(), cmd.KubeconfigCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
