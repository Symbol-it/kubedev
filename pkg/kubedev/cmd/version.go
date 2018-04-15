package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of KubeDev",
		Long:  `All software has versions. This is KubeDev's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("KubeDev kubernetes cluster v0.1 -- HEAD")
		},
	}
}
