package scraper

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Job struct {
	Title    string
	Company  string
	Location string
	Link     string
}

func FetchIndeedJobs(query string) ([]Job, error) {
	searchURL := fmt.Sprintf("https://in.indeed.com/jobs?q=%s", url.QueryEscape(query))

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Full browser-like headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://in.indeed.com/")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Handle gzip if present
	var bodyReader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		bodyReader, err = gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer bodyReader.Close()
	} else {
		bodyReader = res.Body
	}

	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		return nil, err
	}

	var jobs []Job
	doc.Find("h2.jobTitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a span[title]").Text()
		link, _ := s.Find("a").Attr("href")
		company := s.ParentsFiltered("div.job_seen_beacon").Find("span[data-testid='company-name']").Text()
		location := s.ParentsFiltered("div.job_seen_beacon").Find("div[data-testid='text-location'] span").Text()

		jobs = append(jobs, Job{
			Title:    title,
			Company:  company,
			Location: location,
			Link:     "https://in.indeed.com" + link,
		})
	})

	return jobs, nil
}
