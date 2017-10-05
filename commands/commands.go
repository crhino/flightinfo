package commands

import "github.com/concourse/atc"

type JobStats struct {
	Name      string
	Successes int
	Failures  int
}

func BuildFailures(builds []atc.Build) []JobStats {
	return nil
}
