package datetime

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"time"
)

// GetStartDate returns RFC3339 date string
func GetStartDate() (time.Time, error) {
	// start, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:00+00:00")
	start, err := time.Parse(time.RFC3339, "2016-11-07T10:00:00+00:00")

	return start, err
}

// GetValidStartDate parse input string and return time.Time
func GetValidStartDate(start string) time.Time {

	parsed, err := ParseRawTime(start)
	now := time.Now()

	if err != nil {
		log.Warnf("Failed to parse time: %s\n%v+", start, err)
		parsed = now
	} else if now.After(parsed) {
		log.Warnf("Can't get future archive files (%s)", start)
		parsed = now
	}

	return parsed
}

// ParseRawTime parsing string formatted like `2016-11-11T10`
func ParseRawTime(raw string) (time.Time, error) {

	formatted := fmt.Sprintf("%s:00:00+00:00", raw)
	time, err := time.Parse(time.RFC3339, formatted)

	return time, err
}
