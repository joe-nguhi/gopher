package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Learning Go K&D Book
func main() {
	//fmt.Println("Hello World")

	files := os.Args[1:]
	//dup1(files)
	dup1(files)
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
