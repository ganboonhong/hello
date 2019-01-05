package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
	// "net/http/httputil"
	// "reflect"

	"github.com/fatih/color"
)

const (
    hr = "\n------\n\n"
    macmillanEndPoint = "https://dictionaryapi.com/api/v3/references/collegiate/json/"
    wordPause = 700 * time.Millisecond
    sentencePause = 1 * time.Second
)

func getRuntimeGOOS() string {
	return runtime.GOOS
}

func say(s string) {
	if getRuntimeGOOS() != "darwin" {
		return
	}
	cmd := exec.Command("say", s)
	cmd.CombinedOutput()
}

func initHint(s string){

    fmt.Println("\nSearching for " + s + " ...\n")
    say("Searching for");
    time.Sleep(wordPause);
    say(s);
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("No word to search (first arg is the word to search)")
	}

    word := flag.Arg(0)
    initHint(word)

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
                        // color.Blue("Definition " + num + ": ")
                        var totalDefTitle string
                        if len(vv) == 1 {
                            totalDefTitle = "definition"
                        } else {
                            totalDefTitle = "definitions"
                        }
                        resultTitle := fmt.Sprintf(
                            "%d %s \n \n", 
                            len(vv),
                            totalDefTitle,
                        )

                        color.White(hr)

                        color.Blue(resultTitle)
                        say(resultTitle)
                        time.Sleep(wordPause)

						for n, def := range vv {
							num := strconv.Itoa(n + 1)
							color.Blue("Definition " + num + ": ")
							color.Green(def.(string) + "\n\n") // https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion

                            say("Definition " + num)
                            time.Sleep(wordPause)
                            say(def.(string))
                            // say(fmt.Sprintf("Definition %s %s", num, def.(string)))

                            if len(vv) !=  n + 1 {
                                time.Sleep(sentencePause)
                            }
						}

                        color.White(hr)
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
