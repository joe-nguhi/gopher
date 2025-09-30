package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const httpsScheme = "https://"

// SequentialFetch fetches URLs sequentially and saves them to files
func SequentialFetch(links []string) {
	start := time.Now()

	for i, link := range links {
		nstart := time.Now()
		if !strings.HasPrefix(link, httpsScheme) {
			link = fmt.Sprintf("%s%s", httpsScheme, link)
		}

		response, err := http.Get(link)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch:%v\n", err)
			os.Exit(1)
		}

		file, err := os.Create(fmt.Sprintf("generated/link%d.html", i+1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Creating File: %v\n", err)
			os.Exit(1)
		}

		nbytes, err := io.Copy(file, response.Body)

		secs := time.Since(nstart).Seconds()

		fmt.Printf("%.2fs %7d %s \t %s \n", secs, nbytes, link, response.Status)

		response.Body.Close()
		file.Close()
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func ParallelFetch(links []string) {

	start := time.Now()
	ch := make(chan string)

	for i, link := range links {

		if !strings.HasPrefix(link, httpsScheme) {
			link = fmt.Sprintf("%s%s", httpsScheme, link)
		}

		go fetchUrl(link, ch, i)
	}

	for _, _ = range links {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchUrl(link string, ch chan string, index int) {

	start := time.Now()

	response, err := http.Get(link)

	if err != nil {
		ch <- fmt.Sprintf("Fetch:%v\n", err)
	}

	file, err := os.Create(fmt.Sprintf("generated/plink%d.html", index+1))
	if err != nil {
		ch <- fmt.Sprintf("Creating File: %v\n", err)
	}

	nbytes, err := io.Copy(file, response.Body)
	response.Body.Close()
	file.Close()

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s \t %s", secs, nbytes, link, response.Status)

}
