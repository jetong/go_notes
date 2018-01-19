package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/jetong/go_notes/master/http1.go")
	if err != nil {
		fmt.Println(err.Error())
	}else{
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Printf("%s", body)
	}
}
