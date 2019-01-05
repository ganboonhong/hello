package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    // "net/http/httputil"
	"os/exec"
    // "reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
)

const (
	hr                = "\n------\n\n"
	mwEndPoint = "https://dictionaryapi.com/api/v3/references/collegiate/json/"
	wordPause         = 700 * time.Millisecond
	sentencePause     = 1 * time.Second
)

var (
	// mw (merriam-webster) 
    // google (googledictionaryapi) || https://googledictionaryapi.eu-gb.mybluemix.net/ || https://googledictionaryapi.eu-gb.mybluemix.net/?define=computers
	apiProvider = flag.String("agent", "mw", "api provider")
    APIendpointURL string
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

func initHint(s string) {

	fmt.Println("\nSearching for " + s + " ...\n")
	say("Searching for")
	time.Sleep(wordPause)
	say(s)
}

func printDef(defs []string) {

	var totalDefTitle string
	if len(defs) == 1 {
		totalDefTitle = "definition"
	} else {
		totalDefTitle = "definitions"
	}
	resultTitle := fmt.Sprintf(
		"%d %s \n \n",
		len(defs),
		totalDefTitle,
	)

	color.White(hr)

	color.Blue(resultTitle)
	say(resultTitle)
	time.Sleep(wordPause)

	for n, def := range defs {
		num := strconv.Itoa(n + 1)
		color.Blue("Definition " + num + ": ")
		color.Green(def + "\n\n")

		say("Definition " + num)
		time.Sleep(wordPause)
		say(def)
		// say(fmt.Sprintf("Definition %s %s", num, def.(string)))

		if len(defs) != n+1 {
			time.Sleep(sentencePause)
		}
	}

	color.White(hr)

}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("No word to search (first arg is the word to search)")
	}

	word := flag.Arg(0)
	initHint(word)

    switch *apiProvider {
        case "mw":
            APIendpointURL = mwEndPoint + word + "?key=6dfc3570-8a8b-4e4d-8734-aface0fbc277"
    }

	resp, err := http.Get(APIendpointURL)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

    // dump, err := httputil.DumpResponse(resp, true)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Printf("%q", dump)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}


	var f interface{}
	err = json.Unmarshal(bodyBytes, &f)
	if err != nil {
		log.Fatal(err)
	}

    // In this way you can work with unknown JSON data while still enjoying the benefits of type safety.
	results := f.([]interface{}) // https://tour.golang.org/methods/15

	switch *apiProvider {
    case "mw": // merriam-webster api result format: https://dictionaryapi.com/products/api-collegiate-dictionary

    for i, v := range results {            
    	if i == 0 { // use first result
            switch v.(type) {
            case string:
                // target word NOT found (eg. iphone)
                color.Blue("Did you mean: ")
                for _, v := range results {
                    color.White(v.(string));
                }
            case map[string]interface{}:
                // target word found (eg. computer)
                m := v.(map[string]interface{}) // https://blog.golang.org/json-and-go
                for k, v := range m {
                    switch vv := v.(type) { // interface type assertion, https://tour.golang.org/methods/15
                    case []interface{}: // an array type
                    defs := make([]string, len(vv))
                        if k == "shortdef" {
                            for defK, def := range vv {
                                defs[defK] = def.(string) // https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
                            }
                            printDef(defs)
                        }
                    }
                }
             }
    	}
    }
	}

}
