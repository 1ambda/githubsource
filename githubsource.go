package main

import (
	"github.com/urfave/cli"
	"os"

	log "github.com/inconshreveable/log15"

	"github.com/1ambda/githubsource/archive"
	"github.com/1ambda/githubsource/datetime"
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
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "enable dryrun or not",
		},
	}

	app.Action = func(c *cli.Context) error {
		appLog := log.New()
		output := getValidOutput(c.String("output"))
		dryrun := c.Bool("dryrun")
		start := c.String("start")
		end := c.String("end")

		startTime := datetime.GetStartTime(start)
		endTime := datetime.GetEndTime(end)

		if startTime.After(endTime) {
			appLog.Error("Invalid time period", log.Ctx{
				"end":   endTime,
				"start": startTime,
			})
			return nil
		}

		appLog.Info("Execute dryrun", log.Ctx{
			"end":    endTime,
			"start":  startTime,
			"output": output,
			"dryrun": dryrun,
		})

		archive.Download(dryrun, startTime, endTime)

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
