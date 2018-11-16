package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	fmt.Println(b, b["name"])
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":9090")

}
