package main

import (
	"fmt"
	"time"
)

func main() {
	start, _ := getStartDate()
	end, _ := getEndDate()

	for start.Before(end) {
		start = start.Add(time.Hour)

		url := fmt.Sprintf("http://data.githubarchive.org/%d-%02d-%02d-%02d.json.gz",
			start.Year(), start.Month(), start.Day(), start.Hour())

		fmt.Println(url)
	}
}

func getEndDate() (time.Time, error) {
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayFormat :=
		fmt.Sprintf("%d-%02d-%02dT23:00:00+00:00",
			yesterday.Year(),
			yesterday.Month(),
			yesterday.Day())

	end, err := time.Parse(time.RFC3339, yesterdayFormat)

	return end, err
}

func getStartDate() (time.Time, error) {
	// start, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:00+00:00")
	start, err := time.Parse(time.RFC3339, "2016-11-06T18:00:00+00:00")

	return start, err
}
