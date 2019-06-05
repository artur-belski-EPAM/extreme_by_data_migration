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
	CrowlTopics("http://extreme.by/forum/forum28.html")
}

func CrowlTopics(url string) {
	log.Println("Start")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".topictitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("i").Text()
		fmt.Printf("Topic %d: %s\n", i, title)
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")

}
