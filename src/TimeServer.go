package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// handlers
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p align=\"center\">Address: %s%s</p>", r.Host, r.URL.Path)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format(time.RFC1123)
	body := "Current time:"
	fmt.Fprintf(w, "<h1 align=\"center\"> %s </h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\"> %s </h2>", time)
	fmt.Fprintf(w, "<p align=\"center\">Address: %s%s</p>", r.Host, r.URL.Path)

}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 align=\"center\">About</h1>")
	fmt.Fprintf(w, "<h2 align=\"center\">Simple time web server in golang</h2>")
	fmt.Fprintf(w, "<p align=\"center\">Address: %s%s</p>", r.Host, r.URL.Path)
}

func main() {
	fmt.Println("Simple web time server in golang\nby oni-engineer | v1.0")
	var PORT string = ":8080"
	if len(os.Args) == 2 {
		PORT = ":" + os.Args[1]
	}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
