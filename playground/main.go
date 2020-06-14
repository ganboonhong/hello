package main

import(
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type job struct {
	title string
	URL string
}

func getAnchors(url string, keyword string) ([] job) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed getting http response: %v", err)
	}
	defer resp.Body.Close()

	var jobs []job
	tokenizer := html.NewTokenizer(resp.Body)

	// https://golang.org/ref/spec#Break_statements
	L:
		for {
			tokenType := tokenizer.Next()

			switch tokenType {
			case html.ErrorToken:
				err := tokenizer.Err()
				if err != nil {
					if err != io.EOF {
						log.Fatalf("error tokenizing HTML %v", err)
					}
					break L
				}

			case html.StartTagToken:
				token := tokenizer.Token()
				if "a" == token.Data {
					tokenType = tokenizer.Next()
					if tokenType == html.TextToken {
						foundJob := job{}
						for _, a := range token.Attr {
							if a.Key == "href" {
								foundJob.URL = a.Val
								break
							}
						}
						text := tokenizer.Token().Data
						if strings.Contains(strings.ToUpper(text), strings.ToUpper(keyword)) {
							foundJob.title = text 
							jobs = append(jobs, foundJob)
						}
					}
				}
			}
		}

	return jobs
}

func main(){

	if len(os.Args) < 2 {
		log.Fatal("Missing first param")
	}

	keyword := os.Args[1]
	site := "https://www.104.com.tw/jobs/search/?keyword=" + keyword
	jobs := getAnchors(site, keyword)
	
	for _, job := range jobs {
		fmt.Println(job.title, job.URL)	
	}

}