/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"gohire/scraper"
)

// gohireCmd represents the gohire command
var gohireCmd = &cobra.Command{
	Use:   "gohire",
	Short: "A command to search for jobs from multiple platform at one place",
	Long: `This command helps a user search for a job available at different platform like indeed, hirist, etc, at the same time and shows the result here, hence a user dont have to visit multiple websites for same job role.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobTitle, _ := cmd.Flags().GetString("title")
		jobLocation, _ := cmd.Flags().GetString("location")
		jobType, _ := cmd.Flags().GetString("type")

		//Building query string for indeed
		query := jobTitle
		if jobLocation != "" {
			query += " " + jobLocation
		}
		if jobType != ""{
			query += " " + jobType
		}
		
		// fetching jobs from indeed 
		jobs, err := scraper.FetchIndeedJobs(strings.TrimSpace(query))
		if err != nil {
			log.Fatalf("Error fetching jobs: %v", err)
		}

		// Display results
		if len(jobs) == 0 {
			fmt.Println("No jobs found.")
			return
		}

		for i, job := range jobs {
			fmt.Printf("%d. %s | %s | %s\n   %s\n\n", i+1, job.Title, job.Company, job.Location, job.Link)
		} 


	},
}


func init() {
	rootCmd.AddCommand(gohireCmd)

	// Here you will define your flags and configuration settings.
	gohireCmd.PersistentFlags().String("title", "", "A title for job search")
	gohireCmd.PersistentFlags().String("location", "", "Location of the job")
	gohireCmd.PersistentFlags().String("type", "", "Internship, part time or full time")



	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gohireCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gohireCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
