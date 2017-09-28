package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crhino/flightstats/api"
	"github.com/crhino/flightstats/commands"
)

func main() {
	res, err := api.Builds("ci", "https://diego.ci.cf-app.com")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", commands.BuildFailures(res))
}
