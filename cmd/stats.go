package cmd

import (
	"fmt"

	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/hekike/node-report-analytics/pkg/report"
	"github.com/hekike/node-report-analytics/pkg/terminal"
	"github.com/spf13/cobra"
)

// InitStats initialized command
func InitStats(rootCmd *cobra.Command) {
	var dir string
	var cmd = &cobra.Command{
		Use:   "stats",
		Short: "Print the stats of Diagnostics Report(s) from a folder",
		Long:  "Collects all Diagnostics Reports from a specific folder and aggregates them by hash.",
		Run: func(cmd *cobra.Command, args []string) {
			if dir == "" {
				dir = "./"
			}
			reports, err := report.ReadDir(dir)
			if err != nil {
				panic(err)
			}
			if len(reports) == 0 {
				fmt.Printf("No report found in %s\n", dir)
				return
			}

			aggregatedReports := model.AggregateDiagnosticReports(reports)
			terminal.RenderAggregatedDiagnosticReport(aggregatedReports)
		},
	}

	cmd.Flags().StringVarP(
		&dir,
		"dir",
		"d",
		"",
		"directory to scan for reports",
	)

	rootCmd.AddCommand(cmd)
}
