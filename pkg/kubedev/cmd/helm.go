package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewHelmCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "helm",
		Short: "install helm on your GKE cluster",
		Long:  `install helm on your GKE cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			helm := "helm"
			helmArgs := []string{"init"}

			if err := exec.Command(helm, helmArgs...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println("Helm installed on your GKE cluster")
		},
	}
}
