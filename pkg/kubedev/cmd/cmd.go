package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubdev",
	Short: "Kubdev help you to manage your dev cluster",
	Long:  `Kubedev allow you to easily create and delete a dev kubernetes cluster on GKE`,
	Run:   runHelp,
}

func NewKubeDevCmd() *cobra.Command {
	rootCmd.AddCommand(NewVersionCmd())
	rootCmd.AddCommand(NewCreateCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewHelmCmd())
	return rootCmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
