// serve static files in dir static/

package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
