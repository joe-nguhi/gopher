package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpsScheme = "https://"

// SequentialFetch fetches URLs sequentially and saves them to files
func SequentialFetch(links []string) {
	for i, link := range links {
		if !strings.HasPrefix(link, httpsScheme) {
			link = fmt.Sprintf("%s%s", httpsScheme, link)
		}

		response, err := http.Get(link)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch:%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Link: %s:\t Status:%s\n", link, response.Status)

		file, err := os.Create(fmt.Sprintf("generated/link%d.html", i+1))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Creating File: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(file, response.Body)
		response.Body.Close()
		file.Close()
	}
}