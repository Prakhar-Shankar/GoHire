package cmd

import (
	"fmt"
	"gohire/scraper"
	"strings"

	"github.com/spf13/cobra"
)

var (
	titleFilter string
	locationFilter string 
	positionFilter string 
)

var gohireCmd = &cobra.Command{
	Use:   "gohire",
	Short: "Fetch jobs from RemoteOK",
	RunE: func(cmd *cobra.Command, args []string) error {
		jobs, err := scraper.FetchRemoteOKJobs()
		if err != nil {
			return err
		}

		//filtering
		filtered := []scraper.Job{}
		for _, job := range jobs{
			if titleFilter != "" && !strings.Contains(strings.ToLower(job.Position), strings.ToLower(titleFilter)){
				continue
			}
			if locationFilter != "" {
				match := false
				for _, loc := range strings.Split(job.Location, ",") {
					if strings.EqualFold(strings.TrimSpace(loc), locationFilter) {
						match = true
						break
					}
				}
				if !match {
					continue
				}
			}
			if positionFilter != "" && !strings.Contains(strings.ToLower(job.Company), strings.ToLower(positionFilter)) {
				continue
			}
			filtered = append(filtered, job)
		}


		for _, job := range jobs {
			fmt.Printf("%s — %s (%s)\n%s\n\n", job.Company, job.Position, job.Location, job.URL)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gohireCmd)

	gohireCmd.Flags().StringVarP(&titleFilter, "title", "t", "", "Filter jobs by title/position keyword")
	gohireCmd.Flags().StringVarP(&locationFilter, "location", "l", "", "Filter jobs by location")
	gohireCmd.Flags().StringVarP(&positionFilter, "company", "c", "", "Filter jobs by company name")
}
