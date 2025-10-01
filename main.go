package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Learning Go K&D Book

var mu sync.Mutex
var count int

func main() {

	// find duplicates
	/*
		files := os.Args[1:]
		//uniq.Dup1(files)
		uniq.Dup2(files)
	*/

	//Lissajous Figures
	/*
		file, err := os.Create("assets/lissajous.gif")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		animatedgif.Lissajous(file)
	*/

	// Fetch Url
	/*
		links := os.Args[1:]
		fetch.SequentialFetch(links)
		fetch.ParallelFetch(links)
	*/

	// Simple Web Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	port := "8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Path: %s\n", r.URL.Path)
	mu.Lock()
	count++
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d", count)
	mu.Unlock()
}
