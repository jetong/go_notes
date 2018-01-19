package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/jetong/go_notes/master/http1.go")
	if err != nil {
		fmt.Println(err.Error())
	}else{
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s", body)
	}
}
