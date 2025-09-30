package uniq

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Dup1 Unix uniq
// Dup1 Find duplicate lines in given files -v1
func Dup1(files []string) {
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

// Dup2 Find duplicate lines in given files -v2
func Dup2(files []string) {
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
