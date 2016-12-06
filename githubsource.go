package main

import (
	"os"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/urfave/cli"

	"github.com/1ambda/githubsource/archive"
	"github.com/1ambda/githubsource/datetime"
)

func main() {
	app := cli.NewApp()
	app.Name = "githubsource"
	app.Version = "0.0.1"
	app.Usage = "--concurrent -output json --start 2016-11-01T09 --end 2016-11-07T23"
	app.Flags = createFlag()
	app.Action = createAction()
	app.Run(os.Args)
}

func getValidOutput(output string) string {
	if output != "json" {
		return "gz"
	}

	return output
}

func createFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "gz",
			Usage: "file `type`: [gz | json]",
		},
		cli.StringFlag{
			Name:  "start, s",
			Usage: "start `date` (e.g. '2016-11-01T00')",
		},
		cli.BoolFlag{
			Name:  "concurrent, c",
			Usage: "whether to use goroutines or not",
		},
		cli.StringFlag{
			Name:  "end, e",
			Usage: "end `date` (e.g. '2016-11-01T23')",
		},
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "enable dryrun or not",
		},
	}
}

func createAction() func(*cli.Context) error {
	return func(c *cli.Context) error {
		output := getValidOutput(c.String("output"))
		dryrun := c.Bool("dryrun")
		concurrent := c.Bool("concurrent")
		start := c.String("start")
		end := c.String("end")

		startTime := datetime.GetStartTime(start)
		endTime := datetime.GetEndTime(end)

		if startTime.After(endTime) {
			log.Error("Invalid time period", log.Ctx{
				"end":   endTime,
				"start": startTime,
			})
			return nil
		}

		log.Info("Configuration", log.Ctx{
			"end":        endTime,
			"start":      startTime,
			"output":     output,
			"concurrent": concurrent,
			"dryrun":     dryrun,
		})

		now := time.Now()

		archive.Download(concurrent, dryrun, output, startTime, endTime)

		log.Info("Elasped " + time.Since(now).String())

		return nil
	}
}
