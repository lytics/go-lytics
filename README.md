#Lytics SDK for Go
The Lytics SDK for go offers easy integration with all public REST API endpoints. This library is actively being managed and every effort will be made to ensure that all handling reflects the best methods available. Overview of supported methods outlined below.

## Version
0.0.1

## Full REST API Documentation
https://www.getlytics.com/developers/rest-api

## Roadmap
* Command line tool
* Data upload (single and bulk)
* AppEngine specific docs
* More detailed examples

## Getting Started
1. Import the library. `go get github.com/lytics/go-lytics`
2. Create a new client from api key.
3. Run one of the many methods to access account info.

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

## Examples
* [Get All Accounts](examples/get_accounts.md)
* [Get All Segments](examples/get_segments.md)
* [Get Segments A User Belongs To](examples/get_segments_for_user.md)
* [Page Through Users in Segment](examples/page_through_segment.md)

## Supported Methods
* **Account**
	* Single `GET`
	* All `GET`
* **Auth**
	* Single `GET`
	* All `GET`
* **User**
	* Single `GET`
	* All `GET`
* **Provider**
	* Single `GET`
	* All `GET`
* **Work**
	* Single `GET`
	* All `GET` 	
* **Segment**
	* Single `GET`
	* All `GET` 
* **Entity API** `GET`
* **Catalog**
	* Schema `GET`

## Command Line Tool
We have built out a simple command line tool that lets you test many of the endpoints as well as write the results to files. This is a work in progress and will continue to evolve.

### Installation

```
go install github.com/lytics/go-lytics/cmd/lytics
```

### Usage

```
lyticscmd -help
```
	
## Contributing
Want to add something? Go for it, just fork the repo and send us a PR. Please make sure all tests run `go test -v` and that all new functionality comes with well documented and thorough testing.

## License
[Apache Version 2.0 ](LICENSE.md)   
Copyright (c) 2015 Lytics