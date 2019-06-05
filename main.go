package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	CrowlTopics("http://extreme.by/forum/forum28.html")
	// CrowlTopics("http://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks")
}

func CrowlTopics(url string) {
	log.Println("Start")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Get request error:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Parsing error:", err)
	}
	log.Println("Parsing done")
	topics := doc.Find(".topictitle")
	fmt.Printf("found topics: %d \n", topics.Size())
	topics.Each(func(i int, s *goquery.Selection) {
		title := s.Find("i").Text()
		fmt.Printf("==================>Topic %d: %s\n", i, title)
	})

	log.Println("Done")

}
