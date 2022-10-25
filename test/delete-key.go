// CB-Store is a common repository for managing Meta Info of Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
//
// by powerkim@etri.re.kr, 2019.08.

package main

import (
	"fmt"
	"os"
	"strings"

	cbstore "github.com/cloud-barista/cb-store"
	icbs "github.com/cloud-barista/cb-store/interfaces"
)

var store icbs.Store

func init() {
	store = cbstore.GetStore()
}

func main() {

	var key string
	for _, arg := range os.Args {
		if strings.Contains(arg, "=") {
			key = strings.Split(arg, "=")[1]
		}
	}

	fmt.Println("===========================")

	if key != "" {
		// ## Delete
		err := store.Delete(strings.TrimSpace(key))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(">>>> Deleted: " + key)
	} else  {
		fmt.Println(">>>> Can't Delete: " + key)
		fmt.Println("\tExamples)")
		fmt.Println("\t\t$ go run delete-key.go key=/resource-info-spaces/iids:nodegroup/alibaba-tokyo-config/myk8scluser-01/Economy")
	}

	fmt.Println("===========================")
}
