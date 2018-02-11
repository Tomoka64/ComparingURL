package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("<Usage> %v URL1 URL2 word\n", os.Args[0])
		os.Exit(0)
	}
	url1 := os.Args[1]
	url2 := os.Args[2]
	process(url1, "url1.txt")
	process(url2, "url2.txt")
	word := os.Args[3]
	resultForURL1 := WordCount(word, "url1.txt")
	resultForURL2 := WordCount(word, "url2.txt")
	color.Red("Searched the word %v\nResult is-------:\n", word)
	fmt.Printf("URL1(%v) contains %v\n", url1, resultForURL1)
	fmt.Printf("URL2(%v) contains %v\n", url2, resultForURL2)

}

//process makes a file and put the data from URL into the file.
func process(url string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln("could not create a file")
	}
	defer f.Close()

	src, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	content, _ := ioutil.ReadAll(src.Body)

	io.Copy(f, bytes.NewReader(content))

}

//WordCount counts the number of word(os.Args[3]) in the file created by the 'process'.
func WordCount(word string, filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("could not open a file")
	}
	defer f.Close()

	counts := map[string]int{}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		counts[word]++
	}

	return counts[word]
}
