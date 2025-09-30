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
	/*
		links := os.Args[1:]
		fetchUrls(links)
	*/
}

func fetchUrls(links []string) {
	for _, link := range links {

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
		response.Body.Close();

		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: reading: %s, %v\n", link, err)
			os.Exit(1)
		}*/

		//fmt.Printf("%s\n%s\n", link, bytes)

		// Solution 2
		fmt.Printf("Link: %s:\t Status:%s\n", link, response.Status)
		_, err = io.Copy(os.Stderr, response.Body)
	}
}
