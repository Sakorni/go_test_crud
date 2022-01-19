package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,os.Getenv("message"))
	})
	http.ListenAndServe(":8080", nil)
}
