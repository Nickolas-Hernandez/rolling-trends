// https://www.reddit.com/r/bjj/comments/17fa6o5/what_is_the_best_nogi_guard_in_your_opinion/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var redditUrl = "https://www.reddit.com/r/bjj/comments/17fa6o5/what_is_the_best_nogi_guard_in_your_opinion/"

func main() {

	fmt.Println("🔍 Scraping Reddit thread...")

	// Create a new GET request
	req, err := http.NewRequest("GET", redditUrl, nil)
	if err != nil {
		log.Fatal("❌ Failed to create request:", err)
	}

	// Use a browser-like user agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; RollingTrendsBot/1.0)")

	// Execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("❌ Request failed:", err)
	}
	defer res.Body.Close()

	// Check for unsuccessful HTTP status codes
	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("❌ Request failed with status code: %d\n%s", res.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("❌ Failed to read response body:", err)
	}

	// Write HTML content to a file in data/raw
	filePath := "../data/raw/best-no-gi-guards.html"
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		log.Fatal("❌ Failed to write file:", err)
	}

	fmt.Printf("✅ HTML saved to %s\n", filePath)

	// res, err := http.Get(redditUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// f, err := os.CreateTemp("../../data/raw/", "best-no-gi-guards")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := f.Write(body); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }

}
