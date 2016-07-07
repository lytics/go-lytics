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

	// create the segment scanner
	err = client.CreateScanner()
	if err != nil {
		panic(err)
	}

	// start the paging routine
	err = client.PageMembers(segmentid)
	if err != nil {
		panic(err)
	}

	// handle processing the entities
PagingComplete:
	for {
		select {
		case entities := <-client.Scan.Loader:
			for _, v := range entities {
				fmt.Println(v["email"])
			}

		case shutdown := <-client.Scan.Shutdown:
			if shutdown {
				break PagingComplete
			}
		}
	}

	fmt.Printf("*** COMPLETED SCAN: Loaded %d batches and %d total entities", len(client.Scan.Batches), client.Scan.Total)
}

```