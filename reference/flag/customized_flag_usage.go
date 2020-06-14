package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	name := flag.String("name", "", "you must provide your name as first argument")
	flag.Usage = func(){
		fmt.Printf("%s", 
	`Usage (Customized message):  
	1. go run customized_flag_usage.go -name=Francis 
	2. go run customized_flag_usage.go -name Francis // without "=" 

	Default message:

 	`)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *name == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(*name)
}