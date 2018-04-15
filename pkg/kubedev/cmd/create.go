package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "create a new kubernetes cluster on GKE",
		Long:  `create a new kubernetes cluster on GKE`,
		Run: func(cmd *cobra.Command, args []string) {
			gcloud := "gcloud"
			createArgs := []string{"container", "clusters", "create", "kubdev"}
			getCredentialsArgs := []string{"container", "clusters", "get-credentials", "kubdev", "--zone europe-west1-b", "--project kubclust"}

			if err := exec.Command(gcloud, createArgs...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println("KubeDev cluster up and running")

			if err := exec.Command(gcloud, getCredentialsArgs...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println("KubeDev cluster credentials ok")
		},
	}
}
