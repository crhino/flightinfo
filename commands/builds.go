package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/fly/rc"
	"github.com/concourse/go-concourse/concourse"
)

const timeDateLayout = "2006-01-02@15:04:05-0700"

func Builds(user, password string) ([]string, error) {
	target, err := rc.NewBasicAuthTarget("ci", "https://diego.ci.cf-app.com", "", false, user, password, "", false)
	if err != nil {
		return nil, err
	}

	err = target.Validate()
	if err != nil {
		return nil, err
	}

	count := 10
	page := concourse.Page{Limit: count}

	// team := target.Team()
	client := target.Client()

	var builds []atc.Build
	builds, _, err = client.Builds(page)
	if err != nil {
		return nil, err
	}

	var rangeUntil int
	if count < len(builds) {
		rangeUntil = count
	} else {
		rangeUntil = len(builds)
	}

	var results []string
	for _, b := range builds[:rangeUntil] {
		startTime, endTime, duration := populateTimeCells(time.Unix(b.StartTime, 0), time.Unix(b.EndTime, 0))

		var pipelineJobContents, buildContents string
		if b.PipelineName == "" {
			pipelineJobContents = "one-off"
			buildContents = "n/a"
		} else {
			pipelineJobContents = fmt.Sprintf("%s/%s", b.PipelineName, b.JobName)
			buildContents = b.Name
		}

		statusContents := b.Status
		results = append(results, strings.Join([]string{startTime, endTime, duration, pipelineJobContents, buildContents, statusContents}, ", "))
	}

	return results, nil
}

func populateTimeCells(startTime time.Time, endTime time.Time) (string, string, string) {
	var startTimeStr string
	var endTimeStr string
	var durationStr string

	startTime = startTime.Local()
	endTime = endTime.Local()
	zeroTime := time.Unix(0, 0)

	if startTime == zeroTime {
		startTimeStr = "n/a"
	} else {
		startTimeStr = startTime.Format(timeDateLayout)
	}

	if endTime == zeroTime {
		endTimeStr = "n/a"
		durationStr = fmt.Sprintf("%v+", roundSecondsOffDuration(time.Since(startTime)))
	} else {
		endTimeStr = endTime.Format(timeDateLayout)
		durationStr = endTime.Sub(startTime).String()
	}

	if startTime == zeroTime && endTime == zeroTime {
		durationStr = "n/a"
	}

	return startTimeStr, endTimeStr, durationStr
}

func roundSecondsOffDuration(d time.Duration) time.Duration {
	return d - (d % time.Second)
}
