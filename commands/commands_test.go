package commands_test

import (
	"github.com/concourse/atc"
	"github.com/crhino/flightstats/commands"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildFailures", func() {
	var atcBuilds []atc.Build
	BeforeEach(func() {
		atcBuilds = append(atcBuilds,
			atc.Build{
				TeamName:     "main",
				Name:         "617",
				Status:       "succeeded",
				JobName:      "validate-garden-version",
				PipelineName: "main",
				StartTime:    1507161026,
				EndTime:      1507161104,
			})
		atcBuilds = append(atcBuilds,
			atc.Build{
				TeamName:     "main",
				Name:         "664",
				Status:       "started",
				JobName:      "inigo-postgres",
				PipelineName: "main",
				StartTime:    1507161026,
				EndTime:      0,
			})
		atcBuilds = append(atcBuilds,
			atc.Build{
				TeamName:     "main",
				Name:         "663",
				Status:       "failed",
				JobName:      "inigo-mysql",
				PipelineName: "main",
				StartTime:    1507161026,
				EndTime:      1507161027,
			})
		atcBuilds = append(atcBuilds,
			atc.Build{
				TeamName:     "main",
				Name:         "407",
				Status:       "started",
				JobName:      "units-windows",
				PipelineName: "main",
				StartTime:    1507161026,
				EndTime:      0,
			})
		atcBuilds = append(atcBuilds,
			atc.Build{
				TeamName:     "main",
				Name:         "552",
				Status:       "pending",
				JobName:      "units-postgres",
				PipelineName: "main",
				StartTime:    0,
				EndTime:      0,
			})
	})

	It("returns a collection of statistics representing an aggregate history of the pipeline", func() {
		stats := commands.BuildFailures(atcBuilds)
		Expect(stats).NotTo(BeEmpty())
	})
})
