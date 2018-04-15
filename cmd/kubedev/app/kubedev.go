package app

import "github.com/symbol-it/kubedev/pkg/kubedev/cmd"

func Run() error {
	return cmd.NewKubeDevCmd().Execute()
}
