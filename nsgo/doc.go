//Package netskope is an API client developed in GO for use with the Netskope v2 APIs.
//The client was developed for and is used by the (Netskope Terraform Provider) https://github.com/ns-sbrown/terraform-provider-netskope.
//
//Requirements
//
//-	(Go) https://golang.org/doc/install >= 1.17
//
//Usage
//
//	package main
//
//	import (
//		"fmt"
//		"os"
//
//		"github.com/netskopeoss/netskope-api-client-go/netskope"
//	)
//
//	func main() {
//		//Init a client instance
//		nsclient := netskope.NewClient(os.Getenv("NS_BaseURL"), os.Getenv("NS_ApiToken"))
//
//		//Get Publishers
//		pubs, err := nsclient.GetPublishers()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		fmt.Println(pubs)
//
//	}
package nsgo
