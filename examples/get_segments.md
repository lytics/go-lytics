# Get All Segments for Key
Get all of the available segments for a given API key.

```
package main

import (
	"fmt"
	lytics "github.com/lytics/go-lytics"
)

func main() {
	// set your api key
	key := "<YOUR API KEY>"

	// create the client
	client := lytics.NewLytics(key, nil)

	// list all accounts for key
	segments, err := client.GetSegments()
	if err != nil {
		panic(err)
	}

	for _, seg := range segments {
		fmt.Println(seg.Name)
	}
}

```