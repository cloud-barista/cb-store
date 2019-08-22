// CB-Store is a common repository for managing Meta Info of Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
//
// by powerkim@etri.re.kr, 2019.08.

package main

import (
        "fmt"

	"github.com/sirupsen/logrus"
	"github.com/cloud-barista/poc-cb-store/config"
        "github.com/cloud-barista/poc-cb-store"
	icbs "github.com/cloud-barista/poc-cb-store/interfaces"
)

var cblog *logrus.Logger
var store icbs.Store

func init() {
        cblog = config.Cblogger
	store = cbstore.GetStore()
}

func main() {

        cblog.Info("start test!!")

	// ## Test Data & Specs
	keyValueData := []icbs.KeyValue {
		{"/", "root"},
		{"/key1", "value"},
		{"/key1", "value1"},      // same key
		{"/key1/", "value2"},     // end with '/'
		{"/key1/%", "value3%"},   // with special char.
		{"/key1/key2/key3", "value4"}, 
		{"/space key", "space value5"}, 
		{"/newline \n key", "newline \n value6"}, 
		{"/a/b/c/123/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u", "value/value/value"},
	}

	// ### Put
	for _, ev := range keyValueData {
		err := store.Put(ev.Key, ev.Value)
		if err != nil {
			cblog.Error(err)
		}
	}

	// ## Get 
	keyValue, _ := store.Get("/") // exact match

	fmt.Println("=========================== Get(\"/\")")
	fmt.Println("<" + keyValue.Key + "> " + keyValue.Value)
	fmt.Println("===========================")

        keyValue, _ = store.Get("/space key") // exact match

        fmt.Println("=========================== Get(\"space key\")")
        fmt.Println("<" + keyValue.Key + "> " + keyValue.Value)
        fmt.Println("===========================")



        // ## GetAll
        keyValueList, _ := store.GetAll("/", true) // true = Ascending

        fmt.Println("=========================== GetAll(\"/\", Ascending)")
        for _, ev := range keyValueList {
                fmt.Println("<" + ev.Key + "> " + ev.Value)
        }
        fmt.Println("===========================")

        // ## GetAll
        keyValueList, _ = store.GetAll("/", false) // false = Descending

        fmt.Println("=========================== GetAll(\"/\", Descending)")
        for _, ev := range keyValueList {
                fmt.Println("<" + ev.Key + "> " + ev.Value)
        }
        fmt.Println("===========================")

	// ## Delete
	for _, ev := range keyValueData {
		err := store.Delete(ev.Key)
		if err != nil {
			cblog.Error(err)
		}
	}

        cblog.Info("finish test!!")
}
