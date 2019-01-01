package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "strconv"
	// "net/http/httputil"
	// "reflect"

	"github.com/fatih/color"
)

const macmillanEndPoint = "https://dictionaryapi.com/api/v3/references/collegiate/json/"

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("No word to search (first arg is the word to search)")
	}

	word := flag.Arg(0)

	// resp, err := http.Get("http://localhost/ar-webservice/api/testapi.php")
	resp, err := http.Get(macmillanEndPoint + word + "?key=6dfc3570-8a8b-4e4d-8734-aface0fbc277")
	// resp, err := http.Get("https://dictionaryapi.com/api/v3/references/collegiate/json/test?key=6dfc3570-8a8b-4e4d-8734-aface0fbc277")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var f interface{}
	err = json.Unmarshal(bodyBytes, &f)

	// macmillan api result format: https://dictionaryapi.com/products/api-collegiate-dictionary
	// result is an array with different format in different element
	a := f.([]interface{}) // https://tour.golang.org/methods/15

	for i, v := range a {
		if i == 0 {
			// https://blog.golang.org/json-and-go
			m := v.(map[string]interface{})
			for k, v := range m {
				switch vv := v.(type) { // interface type assertion, https://tour.golang.org/methods/15
				// case string:
				//     fmt.Println(k, "is string", vv)
				case []interface{}: // an array type
					if k == "shortdef" {
						for n, def := range vv {
							num := strconv.Itoa(n + 1)
							color.Blue("Definition " + num + " : ")

							color.White(def.(string) + "\n") // https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
						}
						break
					}
				}
			}
		}
	}

	// dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Printf("%q", dump)
}
