package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	// "github.com/1ambda/github-archive-downloader/archive"
	// "github.com/1ambda/github-archive-downloader/datetime"
)

func main() {
	app := cli.NewApp()
	app.Name = "githubsource"
	app.Version = "0.0.1"
	app.Usage = "-output json --start 2016-11-01T00"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "gz",
			Usage: "file `type`: [gz | json]",
		},
		cli.StringFlag{
			Name:  "start, s",
			Usage: "start `date` (e.g. '2016-11-01T00')",
		},
		cli.StringFlag{
			Name:  "end, e",
			Usage: "end `date` (e.g. '2016-11-01T23')",
		},
	}

	app.Action = func(c *cli.Context) error {

		output := getValidOutput(c.String("output"))

		fmt.Println(output)

		// archive.Download(datetime.GetStartDate(), datetime.GetEnddate())
		return nil
	}

	app.Run(os.Args)
}

func getValidOutput(output string) string {
	if output != "json" && output != "gz" {
		return "json"
	}

	return output
}
