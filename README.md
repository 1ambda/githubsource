# githubsource

Download [github archive files](https://www.githubarchive.org/) using CLI.

## Install

```sh
$ go get github.com/1ambda/githubsource
```

## Usage

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

## TODO

- [x] pkg/error
- [x] logger
- [x] decompress gzip
- [x] small test
- [x] CLI command, color
- [x] parallel, sequential options
- [ ] docs 
- [ ] travis
- [ ] badges 
- [ ] packging and distribute as bin
