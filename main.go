package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"nguhi.dev/gopher/animatedgif"
)

const httpsScheme = "https://"

// Learning Go K&D Book
func main() {
	//fmt.Println("Hello World")

	//files := os.Args[1:]
	//dup1(files)
	//dup1(files)

	//Lissajous Figures
	file, err := os.Create("lissajous.gif")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	animatedgif.Lissajous(file)

	// Fetch Url
	/*	links := os.Args[1:]

		fetchUrls(links)*/
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

// Unix uniq
// dup1 Find duplicate lines in given files -v1
func dup1(files []string) {
	stats := make(map[string]map[string]int)

	for _, file := range files {
		c := make(map[string]int)

		f, err := os.Open(file)
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			c[scanner.Text()]++
		}

		stats[file] = c

		f.Close()
	}

	printResults(stats)
}

// dup2 Find duplicate lines in given files -v2
func dup2(files []string) {
	stats := make(map[string]map[string]int)

	for _, file := range files {
		c := make(map[string]int)

		bytes, err := os.ReadFile(file)

		if err != nil {
			continue
		}

		lines := strings.Split(string(bytes), "\n")

		for _, line := range lines {
			c[line]++
		}

		stats[file] = c
	}

	printResults(stats)
}

func printResults(stats map[string]map[string]int) {
	fmt.Println("------------ Duplicates Summary -------------")

	for name, counter := range stats {
		fmt.Printf("Filename: \t%s\n", name)
		fmt.Println("Count: \t Line")
		fmt.Println(strings.Repeat("---", 10))
		for line, c := range counter {
			if c > 1 {
				fmt.Printf("%d\t%s\n", c, line)
			}
		}
	}
}
