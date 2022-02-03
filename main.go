// TODO: ADD LANGUAGE SELECTION


package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"flag"
)

const URL = "https://tureng.com/tr/turkce-ingilizce/"


func main() {

	// FLAG OPERATIONS
	var word string
	var ret int

	flag.StringVar(&word, "w", "", "word for searching")
	flag.IntVar(&ret, "ret", 1, "return word length")
	flag.Parse()

	if word == "" {
		log.Fatalln("Error: Word Parameter must not be empty\n]More Information for --help flag")
		os.Exit(1)
	}
	
	resp, err := http.Get(URL+word)

	if err != nil {
		log.Fatalln(err)
	}

	
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	words := []string{}

	doc.Find(".ts").Each(func(i int, s *goquery.Selection) {
		response := s.Find("a").Text()

		words = append(words, response)
	})

	for i:= 0; i< ret; i++ {
		fmt.Println(i+1, ": "+words[i])
	}

	fmt.Println("\nDone :D")
}
