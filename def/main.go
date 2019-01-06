package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "net/http/httputil" // DumpResponse
	"os/exec"
	// "reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
)

const (
	hr             = "------\n\n"
	mwEndPoint     = "https://dictionaryapi.com/api/v3/references/collegiate/json/"
	googleEndPoint = "https://googledictionaryapi.eu-gb.mybluemix.net/?define="
	// googleEndPoint = "https://mydictionaryapi.appspot.com/?define="
	wordPause     = 700 * time.Millisecond
	sentencePause = 1 * time.Second
)

var (
	apiProvider    = flag.String("a", "google", "api provider") // https://medium.com/@martin.breuss/finding-a-useful-dictionary-api-52084a01503d
	wSpeech        = flag.Bool("s", false, "read out the definitions")
	APIendpointURL string
)

type GDefinition []struct {
	Definition string   `json:definition,omitempty`
	Example    string   `json:example,omitempty`
	Synonyms   []string `json:example,omitempty`
}

// nested struct
type GResponse []struct {
	Word     string `json:word,omitempty`
	Phonetic string `json:phonetic,omitempty`
	Meaning  struct {
		Noun   GDefinition `json:noun,omitempty`
		Verb   GDefinition `json:verb,omitempty`
		Adverb GDefinition `json:adverb,omitempty`
	} `json:meaning,omitempty`
}

func getRuntimeGOOS() string {
	return runtime.GOOS
}

func say(s string) {
	if getRuntimeGOOS() != "darwin" || *wSpeech == false {
		return
	}
	cmd := exec.Command("say", s)
	cmd.CombinedOutput()
}

func initHint(s string) {

	fmt.Println(
		fmt.Sprintf("\nSearching for %q ...", s),
	)
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
		"%d %s %s\n \n",
		len(defs),
		totalDefTitle,
		"found:",
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

		if len(defs) != n+1 {
			time.Sleep(sentencePause)
		}
	}
	color.White(hr)
}

func gSetDefs(a GDefinition, defs []string) []string {
	for _, v := range a {
		if v.Definition != "" {
			defs = append(defs, v.Definition)
		}
	}
	return defs
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
	case "google":
		APIendpointURL = googleEndPoint + word

	}
	resp, err := http.Get(APIendpointURL)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	switch *apiProvider {
	case "mw":
		var f interface{}
		err = json.Unmarshal(bodyBytes, &f)

		// In this way you can work with unknown JSON data while still enjoying the benefits of type safety.
		results := f.([]interface{}) // https://tour.golang.org/methods/15

		if err != nil {
			log.Fatal(err)
		}
		// merriam-webster
		// format: https://dictionaryapi.com/products/api-collegiate-dictionary
		// url: https://dictionaryapi.com/api/v3/references/collegiate/json/computer?key=6dfc3570-8a8b-4e4d-8734-aface0fbc277
		for i, v := range results {
			if i == 0 { // use first result
				switch v.(type) {
				case string:
					// target word NOT found (eg. iphone)
					color.Blue("Did you mean: ")
					for _, v := range results {
						color.White(v.(string))
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

	case "google":
		// github: https://github.com/meetDeveloper/googleDictionaryAPI
		// endpoint 1: https://mydictionaryapi.appspot.com/
		// endpoint 2: https://googledictionaryapi.eu-gb.mybluemix.net
		// example url: https://googledictionaryapi.eu-gb.mybluemix.net/?define=computer

		var gResponse GResponse
		json.Unmarshal(bodyBytes, &gResponse)
		defs := make([]string, 0)
		for _, r := range gResponse {
			if len(r.Meaning.Noun) > 0 {
				defs = gSetDefs(r.Meaning.Noun, defs)
			}

			if len(r.Meaning.Verb) > 0 {
				defs = gSetDefs(r.Meaning.Verb, defs)
			}

			if len(r.Meaning.Adverb) > 0 {
				defs = gSetDefs(r.Meaning.Adverb, defs)
			}
		}
		printDef(defs)
	}

}
