package scraper

import (
	"encoding/json"
	"net/http"
	"time"
)

type Job struct {
	Company  string `json:"company"`
	Position string `json:"position"`
	Location string `json:"location"`
	URL      string `json:"url"`
}

func FetchRemoteOKJobs() ([]Job, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://remoteok.com/api")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jobs []Job
	if err := json.NewDecoder(resp.Body).Decode(&jobs); err != nil {
		return nil, err
	}

	if len(jobs) > 0 {
		jobs = jobs[1:]
	}

	return jobs, nil
}
