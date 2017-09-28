package api

import (
	"github.com/concourse/atc"
	"github.com/concourse/fly/rc"
	"github.com/concourse/go-concourse/concourse"
)

func Builds(targetName, api string) ([]atc.Build, error) {
	target, err := rc.NewNoAuthTarget("ci", "https://diego.ci.cf-app.com", "", false, "", false)
	if err != nil {
		return nil, err
	}

	err = target.Validate()
	if err != nil {
		return nil, err
	}

	count := 10000
	page := concourse.Page{Limit: count}

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

	return builds[:rangeUntil], nil
}
