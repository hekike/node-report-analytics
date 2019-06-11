package main

import (
	"github.com/hekike/node-report-analytics/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "node-report-analytics"}
	cmd.InitElastic(rootCmd)
	cmd.InitStats(rootCmd)
	rootCmd.Execute()
}
