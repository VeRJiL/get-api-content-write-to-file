package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type response struct {
	data  []byte
	error error
}

func getOneConcurrently(c chan response, wg *sync.WaitGroup, counter int) {
	defer wg.Done()
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", counter)
	res, err := http.Get(url)
	
	if err != nil {
		c <- response{data: nil, error: fmt.Errorf("cant get to the webpage")}
	}
	
	if res.StatusCode != http.StatusOK {
		c <- response{data: nil, error: fmt.Errorf("Skipping Commic : %d, Got: %d\n", counter, res.StatusCode)}
	}
	
	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		c <- response{data: nil, error: fmt.Errorf("Invalid Body: %d\n", err)}
	}
	
	res.Body.Close()
	
	c <- response{data: body, error: nil}
}

func main() {
	start := time.Now()
	
	var fails, counter int
	var data = make(chan response)
	var loopScape = make(chan bool)
	var loopBreak = make(chan bool)
	wg := sync.WaitGroup{}
	
	f, err := os.Create("xkcd.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't Create the file")
	}
	
	defer f.Close()
	
	f.WriteString("[\n")
	
	for i := 1; fails < 2; i++ {
		
		wg.Add(1)
		go getOneConcurrently(data, &wg, i)
		
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			response := <-data
			
			if response.error == nil {
				fails++
				loopScape <- false
			}
			
			if counter > 0 {
				f.WriteString(",\n")
			}
			
			if counter == 100 {
				loopScape <- false
			}
			
			f.Write(response.data)
			
			counter++
			fails = 0
			fmt.Printf("Read comic number: %d\n\n", i)
		}()
		
		if !<-loopScape {
			continue
		}
		
		if !<-loopBreak {
			break
		}
	}
	
	f.WriteString("\n]")
	
	wg.Wait()
	
	secs := time.Since(start).Seconds()
	fmt.Printf("Read Total Number Of %d Comics\n", counter)
	fmt.Printf("It tool: %v seconds", secs)
}
