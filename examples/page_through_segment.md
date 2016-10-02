# Segment Paging
Iterate through all members of a particular segment. Returns the entire model for each user.

```
package main

import (
	"fmt"
	lytics "github.com/lytics/go-lytics"
)

func main() {
	var err error

	// set your api key & segment id to scan
	apikey := "<YOUR API KEY HERE>"
	segmentid := "<YOUR SEGMENT ID HERE>"

	// create the client
	client := lytics.NewLytics(apikey, nil, nil)

	// create the scanner
	scan := client.PageSegmentId(os.Args[1])

	// handle processing the entities
	for {
		e := scan.Next()
		if e == nil {
			break
		}

		fmt.Printf("%v\n\n", e.PrettyJson())
	}

	fmt.Printf("*** COMPLETED SCAN: Loaded %d batches and %d total entities", len(scan.Batches), scan.Total)
}

```