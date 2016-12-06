package archive

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

// Download github archive files
func Download(concurrent, dryrun bool, output string, start, end time.Time) {
	if concurrent {
		concurrentDownload(dryrun, output, start, end)
	} else {
		sequentialDownload(dryrun, output, start, end)
	}
}

// Download github archive files concurrently
func concurrentDownload(dryrun bool, output string, start, end time.Time) {
	var err error
	var wg sync.WaitGroup

	for start.Before(end) {
		start = start.Add(time.Hour)
		prefix := fmt.Sprintf("%d-%02d-%02d-%02d",
			start.Year(), start.Month(), start.Day(), start.Hour())
		url := fmt.Sprintf("http://data.githubarchive.org/%s.json.gz", prefix)
		context := log.Ctx{"prefix": prefix, "output": output}

		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			err = nil
			if !dryrun {
				err = getGithubSource(output, url, prefix)
			}
			if err != nil {
				log.Error(err.Error(), context)
				return
			}
			log.Info(fmt.Sprintf("Downloaded %s.%s", prefix, output), context)
		}()
	}

	wg.Wait()
}

// Download github archive files
func sequentialDownload(dryrun bool, output string, start, end time.Time) {
	var err error

	for start.Before(end) {
		start = start.Add(time.Hour)
		prefix := fmt.Sprintf("%d-%02d-%02d-%02d",
			start.Year(), start.Month(), start.Day(), start.Hour())
		url := fmt.Sprintf("http://data.githubarchive.org/%s.json.gz", prefix)
		context := log.Ctx{"prefix": prefix, "output": output}

		err = nil
		if !dryrun {
			err = getGithubSource(output, url, prefix)
		}
		if err != nil {
			log.Error(err.Error(), context)
			continue
		}
		log.Info(fmt.Sprintf("Downloaded %s.%s", prefix, output), context)
	}
}

func getGithubSource(output, url, prefix string) error {
	// 1. Get json
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		err = errors.Errorf("Failed to get %s (404)", url)
		return err
	}

	if output == "json" {
		filename := fmt.Sprintf("%s.%s", prefix, output)
		err = getJsonSource(res, filename)
	} else {
		filename := fmt.Sprintf("%s.%s", prefix, "gz")
		err = getGzipSource(res, filename)
	}

	return err
}

func getGzipSource(res *http.Response, filename string) error {
	// 3. Write to file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}

func getJsonSource(res *http.Response, filename string) error {
	// 2. Unpack gzip concurrently
	pipeReader, pipeWriter := io.Pipe()

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

	_, err = io.Copy(out, pipeReader)

	return err
}
