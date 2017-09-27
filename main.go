package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crhino/flightstats/commands"
)

func main() {
	user := os.Getenv("CONCOURSE_USER")
	password := os.Getenv("CONCOURSE_PASSWORD")

	res, err := commands.Builds(user, password)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", res)
}
