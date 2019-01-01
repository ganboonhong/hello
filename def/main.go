package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    // "net/http/httputil"
    // "reflect"
    
    "github.com/fatih/color"
)

const macmillanEndPoint = "https://dictionaryapi.com/api/v3/references/sd3/json/"

func main() {
    // ("test", "collegiate", "6dfc3570-8a8b-4e4d-8734-aface0fbc277");
    // ($word, $ref, $key) {
    // $uri = "https://dictionaryapi.com/api/v1/references/" . urlencode($ref) . "/xml/" . urlencode($word) . "?key=" . urlencode($key);
  
    // resp, err := http.Get("http://localhost/ar-webservice/api/testapi.php")
	resp, err := http.Get(macmillanEndPoint + "computer?key=76ffdfcb-e02b-4f16-83c7-8e388a38972d")
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
                        for _, def := range vv {
                            color.Blue("Definition: ")

                            color.White(def.(string)) // https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
                        }
                        break;
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
