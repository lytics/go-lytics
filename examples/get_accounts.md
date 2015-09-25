# Get All Accounts for Key
Get all of the available accounts and child accounts for a given API key.

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
	accounts, err := client.GetAccounts()
	if err != nil {
		panic(err)
	}

	for _, acct := range accounts {
		fmt.Println(acct.Name)
	}
}
```