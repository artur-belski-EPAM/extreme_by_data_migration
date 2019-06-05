package main

import (
	"fmt"
	"net/http"

	//   "net/http"
	// 	"os"
	// 	"strings"
	// 	"golang.org/x/net/html"

	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	log.Print("Start")
	res, err := http.Get("http://m.extreme.by/forum/forum28.html?sid=49ffd60e53d2e8b69988cf46eb48894d")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Print(doc)

	log.Print("Done")

	if err != nil {
		log.Fatal(err)
	}

}
