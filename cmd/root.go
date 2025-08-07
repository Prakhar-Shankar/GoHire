package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "gohire",
    Short: "gohire is a cli tool for searching jobs by scraping job hunting sites.",
    Long:  "gohire is a cli tool for searching jobs by scraping sites like indeed, hirist, etc, to pull all the job data at one place.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
        os.Exit(1)
    }
}