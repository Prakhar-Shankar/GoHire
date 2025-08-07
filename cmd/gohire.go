/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gohireCmd represents the gohire command
var gohireCmd = &cobra.Command{
	Use:   "gohire",
	Short: "A command to search for jobs from multiple platform at one place",
	Long: `This command helps a user search for a job available at different platform like indeed, hirist, etc, at the same time and shows the result here, hence a user dont have to visit multiple websites for same job role.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobTitle, _ := cmd.Flags().GetString("title")
		if jobTitle == "Golang"{
			fmt.Printf("title: %s\n", jobTitle)
		}else{
			fmt.Printf("Wrong title\n")
		}

		jobLocation, _ := cmd.Flags().GetString("location")
		if jobLocation == "Remote"{
			fmt.Printf("location: %s\n", jobLocation)
		}else{
			fmt.Printf("Wrong location\n")
		}

		jobType, _ := cmd.Flags().GetString("type")
		if jobType == "Intern"{
			fmt.Printf("type: %s\n", jobType)
		}else{
			fmt.Printf("Wrong type\n")
		}
	},
}

func getTitle(jobTitle string) {
    fmt.Printf("You searched for a job: %v",  jobTitle)
}

func getLocation(jobLocation string) {
    fmt.Printf("You searched for a job: %v",  jobLocation)
}

func getType(jobType string) {
    fmt.Printf("You searched for a job: %v",  jobType)
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
