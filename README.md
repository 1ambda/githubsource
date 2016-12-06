[![Go Report Card](https://goreportcard.com/badge/github.com/1ambda/githubsource)](https://goreportcard.com/report/github.com/1ambda/githubsource) [![Travis Build](https://travis-ci.org/1ambda/githubsource.svg?branch=master)](https://travis-ci.org/1ambda/githubsource.svg?branch=master)

# githubsource

Download [github archive files](https://www.githubarchive.org/) using CLI.

## Install

```sh
$ go get github.com/1ambda/githubsource
```

## Usage

```
# concurrently download archive files from 2016-11-01T09 until 2016-11-07T23 while extracting gz to json

$ githubsource --concurrent -output json --start 2016-11-01T09 --end 2016-11-07T23"
```

```sh
$ githubsource - h

NAME:
   githubsource - --concurrent -output json --start 2016-11-01T09 --end 2016-11-07T23

USAGE:
   githubsource [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output type, -o type  file type: [gz | json] (default: "gz")
   --start date, -s date   start date (e.g. '2016-11-01T00')
   --concurrent, -c        whether to use goroutines or not
   --end date, -e date     end date (e.g. '2016-11-01T23')
   --dryrun, -d            enable dryrun or not
   --help, -h              show help
   --version, -v           print the version
```
