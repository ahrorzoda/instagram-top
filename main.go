package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	file, err := os.Create("instagram_top_50.csv")
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Rating", "Name", "Username"})
	c.OnHTML(".top-user-details", func(e *colly.HTMLElement) {
		rating := e.ChildText(".rank-num")
		name := e.ChildText(".top-username")
		username := strings.TrimSpace(e.ChildText(".top-nickname"))
		writer.Write([]string{rating, name, username})
	})

	// Start the scraping process
	c.Visit("https://hypeauditor.com/top-instagram-all-russia/")

	fmt.Println("Scraping completed. Data is saved in instagram_top_50.csv")
}
