package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete a kubernetes cluster on GKE",
		Long:  `delete a kubernetes cluster on GKE`,
		Run: func(cmd *cobra.Command, args []string) {
			gcloud := "gcloud"
			arguments := []string{"container", "clusters", "delete", "kubdev", "--quiet"}

			if err := exec.Command(gcloud, arguments...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println("KubeDev deleted")
		},
	}
}
