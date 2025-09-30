package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpsScheme = "https://"

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
	fetchUrls(links)

}

func fetchUrls(links []string) {
	for i, link := range links {

		if !strings.HasPrefix(link, httpsScheme) {
			link = fmt.Sprintf("%s%s", httpsScheme, link)
		}

		response, err := http.Get(link)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch:%v\n", err)
			os.Exit(1)
		}

		// Solution  1
		/*bytes, err := io.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: reading: %s, %v\n", link, err)
			os.Exit(1)
		}

		//fmt.Printf("%d.%s\n%s\n", i, link, bytes)

		file, err := os.Create(fmt.Sprintf("generated/link%d.html", i+1))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Creating File: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(file, string(bytes))*/

		// Solution 2
		fmt.Printf("Link: %s:\t Status:%s\n", link, response.Status)

		file, err := os.Create(fmt.Sprintf("generated/link%d.html", i+1))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Creating File: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(file, response.Body)
	}
}
