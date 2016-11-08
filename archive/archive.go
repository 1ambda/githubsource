package archive

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

// Download github archive files
func Download() {
	start, _ := getStartDate()
	end, _ := getEndDate()
	var err error = nil

	for start.Before(end) {
		start = start.Add(time.Hour)

		filename := fmt.Sprintf("%d-%02d-%02d-%02d.json",
			start.Year(), start.Month(), start.Day(), start.Hour())
		url := fmt.Sprintf("http://data.githubarchive.org/%s.gz", filename)

		err = getGzipJsonAndWriteToFile(url, filename)

		if err == nil {
			log.Info("Saved - ", filename)
		} else {
			log.Error(fmt.Sprintf("%+v\n", err))
		}
	}
}

func getGzipJsonAndWriteToFile(url string, filename string) error {

	// 1. Get json
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		err = errors.Errorf("Failed to get %s (404)", url)
		return err
	}

	// 2. Unpack gzip concurrently
	pipeRedaer, pipeWriter := io.Pipe()

	go func() {
		gzipReader, _ := gzip.NewReader(res.Body)

		defer func() {
			res.Body.Close()
			gzipReader.Close()
			pipeWriter.Close()
		}()

		io.Copy(pipeWriter, gzipReader)
	}()

	// 3. Write to file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, pipeRedaer)

	return err
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
	start, err := time.Parse(time.RFC3339, "2016-11-07T10:00:00+00:00")

	return start, err
}
