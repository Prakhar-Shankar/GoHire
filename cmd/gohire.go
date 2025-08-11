package cmd

import (
	"fmt"
	"gohire/scraper"
	"github.com/spf13/cobra"
)

var gohireCmd = &cobra.Command{
	Use:   "gohire",
	Short: "Fetch jobs from RemoteOK",
	RunE: func(cmd *cobra.Command, args []string) error {
		jobs, err := scraper.FetchRemoteOKJobs()
		if err != nil {
			return err
		}

		for _, job := range jobs {
			fmt.Printf("%s — %s (%s)\n%s\n\n", job.Company, job.Position, job.Location, job.URL)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gohireCmd)
}
