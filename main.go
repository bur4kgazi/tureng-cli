// TODO: ADD LANGUAGE SELECTION

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"tureng/modules"

	"github.com/PuerkitoBio/goquery"
)

var URL = "https://tureng.com/en/"

func main() {

	config := modules.Get("/home/bur4k/.config/tureng/config")

	// FLAG OPERATIONS
	var word string
	var ret int
	var lang string // tr-en de-en es-en fr-en en

	flag.StringVar(&word, "w", "", "word for searching")
	flag.StringVar(&lang, "lang", config.Lang, "which language do you translate")
	flag.IntVar(&ret, "ret", config.Res, "return word length")
	flag.Parse()

	if lang == "tr-en" {
		URL += "turkish-english/"
	} else if lang == "de-en" {
		URL += "german-english/"
	} else if lang == "es-en" {
		URL += "spanish-english/"
	} else if lang == "fr-en" {
		URL += "french-english/"
	} else if lang == "en" {
		URL += "english-synonym/"
	}

	if word == "" {
		fmt.Println("Error: Word Parameter must not be empty")
		flag.Usage()
		os.Exit(1)
	}

	resp, err := http.Get(URL + word)

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

	for i := 0; i < ret; i++ {
		fmt.Println(i+1, ": "+words[i])
	}

	fmt.Println("\nDone :D")
}
