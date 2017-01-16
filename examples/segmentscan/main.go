package main

import (
	"fmt"
	"os"

	lytics "github.com/lytics/go-lytics"
)

func main() {
	lyticsBanner()
	apikey := os.Getenv("LIOKEY")
	if apikey == "" {
		fmt.Println("must have  LIOKEY env set")
		usageExit()
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		usageExit()
	}

	// create the client
	client := lytics.NewLytics(apikey, nil, nil)

	// create the scanner
	scan := client.PageSegment(os.Args[1])

	// handle processing the entities
	for {
		e := scan.Next()
		if e == nil {
			break
		}

		fmt.Println(e.PrettyJson())
	}

	fmt.Printf("*** COMPLETED SCAN: Loaded %d total entities\n\n", scan.Total)
}

func lyticsBanner() {
	fmt.Printf(`
    __          __  _
   / /   __  __/ /_(_)_________
  / /   / / / / __/ / ___/ ___/
 / /___/ /_/ / /_/ / /__(__  )
/_____/\__, /\__/_/\___/____/
      /____/
`)
}
func usageExit() {
	fmt.Printf(`

**************  Page Segment of Users HELP  **************

export LIOKEY="mykey"


# change directory to this segment scan

cd segmentscan

# build this example tool 

go build


segmentscan segment_id

`)
	os.Exit(1)
}
