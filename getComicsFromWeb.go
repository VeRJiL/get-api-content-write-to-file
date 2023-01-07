package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOne(counter int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", counter)
	res, err := http.Get(url)
	
	if err != nil {
		fmt.Fprint(os.Stderr, "Cant get to the webpage")
		os.Exit(-1)
	}
	
	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Skipping Commic : %d, Got: %d", counter, res.StatusCode)
		return nil
	}
	
	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Body: %d", err)
		os.Exit(-1)
	}
	
	res.Body.Close()
	
	return body
}

func main() {
	var fails, counter int
	
	f, err := os.Create("xkcd.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't Create the file")
	}
	
	f.WriteString("[\n")
	
	for i := 1; fails < 2; i++ {
		
		data := getOne(i)
		
		if data == nil {
			fails++
			continue
		}
		
		if counter > 0 {
			f.WriteString(",\n")
		}
		
		f.Write(data)
		
		counter++
		fails = 0
		fmt.Printf("Read comic number: %d\n", i)
	}
	
	f.WriteString("\n]")
	f.Close()
	
	fmt.Printf("Read Total Number Of %d Comics", counter)
}
