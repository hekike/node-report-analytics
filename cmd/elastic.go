package cmd

import (
	"context"
	"fmt"

	"github.com/hekike/node-report-analytics/pkg/db"
	"github.com/hekike/node-report-analytics/pkg/report"
	"github.com/spf13/cobra"
)

// InitElastic initialized command
func InitElastic(rootCmd *cobra.Command) {
	var dir string
	var cmd = &cobra.Command{
		Use:   "elastic",
		Short: "Loads Diagnostics Report(s) to Elasticsearch",
		Long:  `Collects all Diagnostics Reports from a specific folder and loads them to Elasticsearch.`,
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

			ctx := context.Background()

			// init Elastic client
			elasticClient, err := db.NewElasticClient(ctx)
			if err != nil {
				panic(err)
			}

			err = db.CreateReportIndexIfDoesNotExist(ctx, elasticClient)
			if err != nil {
				panic(err)
			}

			err = db.InsertReports(ctx, elasticClient, reports)
			if err != nil {
				panic(err)
			}

			fmt.Println("Documents saved:")
			for _, report := range reports {
				fmt.Printf(
					"%s, %s, %s\n",
					report.Hash,
					report.Header.DumpEventTime,
					report.JavaScriptStack.Message,
				)
			}
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
