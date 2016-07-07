# Get Segment Membership for User
Does a search by email field to find a single user and then print out the segments that they are currently a member of.

```
package main

import (
	"fmt"
	lytics "github.com/lytics/go-lytics"
)

func main() {
	key := "<YOUR API KEY>"            // your primary api key
	entitytype := "<YOUR ENTITY TYPE>" // type of entity to query (e.g. user, content)
	fieldname := "<YOUR FIELD NAME>"   // name of field to query against (e.g. email, _uid)
	fieldval := "<YOUR FIELD VALUE>"   // value of the field

	// create the client
	client := lytics.NewLytics(key, nil, nil)

	// list all accounts for key
	entity, err := client.GetEntity(entitytype, fieldname, fieldval, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(entity["segments"])
}
```