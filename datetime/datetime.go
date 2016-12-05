package datetime

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"time"
)


// GetEndTime parse input string and return time.Time
func GetStartTime(start string) time.Time {
        yesterday := time.Now().AddDate(0, 0, -1)

        if start == "" {
                yesterdayTimeStr := getCurrentTimeStr(yesterday)
                start = yesterdayTimeStr
        }

        parsed, err := ParseRawTime(start)

        if err != nil {
                log.Warnf("Failed to parse time: %s\n%v+", start, err)
                parsed = yesterday
        }

        return parsed
}

// GetEndTime parse input string and return time.Time
func GetEndTime(end string) time.Time {
        now := time.Now()

	if end == "" {
                currentTimeStr := getCurrentTimeStr(now)
		end = currentTimeStr
	}

	parsed, err := ParseRawTime(end)

	if err != nil {
		log.Warnf("Failed to parse time: %s\n%v+", end, err)
		parsed = now
	}

	return parsed
}

// getCurrentTimeStr return current time formatted like `2016-11-11T10`
func getCurrentTimeStr(now time.Time) string {
        return fmt.Sprintf("%04d-%02d-%02dT%02d",
                now.Year(), now.Month(), now.Day(), now.Hour())
}

// ParseRawTime parsing string formatted like `2016-11-11T10`
func ParseRawTime(raw string) (time.Time, error) {
	formatted := fmt.Sprintf("%s:00:00+00:00", raw)
	time, err := time.Parse(time.RFC3339, formatted)

	return time, err
}
