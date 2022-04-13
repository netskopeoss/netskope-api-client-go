# NETSKOPE-API-CLIENT-GO

NETSKOPE-API-CLIENT-GO is an API client developed in GO for use with the Netskope v2 APIs.
The client was developed for and is used by the [Netskope Terraform Provider](https://github.com/netskopeoss/terraform-provider-netskope).

## Requirements

-	[Go](https://golang.org/doc/install) >= 1.17

Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/netskopeoss/netskope-api-client-go/nsgo"
)

func main() {
	//Init a client instance
	nsclient := nsgo.NewClient(os.Getenv("NS_BaseURL"), os.Getenv("NS_ApiToken"))

	//Get Publishers
	pubs, err := nsclient.GetPublishers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pubs)

}
```