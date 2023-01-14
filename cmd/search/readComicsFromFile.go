package search

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Comic struct {
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	Transcript string `json:"transcript"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
}

func (c *Comic) String() {
	fmt.Println("{")
	fmt.Printf("\tNum: %d\n", c.Num)
	fmt.Printf("\tSafeTitle: %s\n", c.SafeTitle)
	fmt.Printf("\tAlt: %s\n", c.Alt)
	fmt.Printf("\tImage: %s\n", c.Image)
	fmt.Printf("\tTtitle: %s\n", c.Title)
	fmt.Printf("\tYear: %s\n", c.Year)
	fmt.Printf("\tDay: %s\n", c.Day)
	fmt.Println("}")
}

func Search(fileName string, inputSearchTerms []string) {
	
	var (
		comics      []Comic
		searchTerms []string
		count       int
		err         error
		f           io.ReadCloser
	)
	
	if f, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Cant Open the File: %s", err)
		os.Exit(-1)
	}
	
	if err = json.NewDecoder(f).Decode(&comics); err != nil {
		fmt.Fprintf(os.Stderr, "Bad Json: %s", err)
		os.Exit(-1)
	}
	
	for _, term := range inputSearchTerms {
		searchTerms = append(searchTerms, strings.ToLower(term))
	}

outer:
	for _, comic := range comics {
		title := strings.ToLower(comic.Title)
		transcript := strings.ToLower(comic.Transcript)
		
		for _, term := range searchTerms {
			if !strings.Contains(title, term) || !strings.Contains(transcript, term) {
				continue outer
			}
		}
		
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %q\n",
			comic.Num,
			comic.Month,
			comic.Day,
			comic.Year,
			comic.Title,
		)
		
		count++
	}
	
	fmt.Printf("Read Total %d Comics\n", len(comics))
	fmt.Printf("Found %d Comics\n", count)
}
