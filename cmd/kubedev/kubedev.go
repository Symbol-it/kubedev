package main

import (
	"os"

	"github.com/symbol-it/kubedev/cmd/kubedev/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
