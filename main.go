package main

import (
	"os"

	"nguhi.dev/gopher/fetch"
)

// Learning Go K&D Book
func main() {

	// find duplicates
	/*
		files := os.Args[1:]
		uniq.Dup1(files)
		uniq.Dup2(files)
	*/

	//Lissajous Figures
	/*
		file, err := os.Create("lissajous.gif")

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		animatedgif.Lissajous(file)
	*/

	// Fetch Url
	links := os.Args[1:]
	fetch.SequentialFetch(links)
	fetch.ParallelFetch(links)
}
