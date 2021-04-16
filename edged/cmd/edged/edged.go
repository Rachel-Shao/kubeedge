package main

import (
	"os"
	"github.com/kubeedge/kubeedge/edged/cmd/edged/app"
	"k8s.io/component-base/logs"
)

func main() {
	command := app.NewEdgedCommand()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}