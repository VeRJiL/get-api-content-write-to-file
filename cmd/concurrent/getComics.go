package concurrent

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

func getOne(c chan response, counter int) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", counter)
	res, err := http.Get(url)
	
	if err != nil {
		c <- response{data: nil, error: fmt.Errorf("cant get to the webpage")}
		return
	}
	
	if res.StatusCode != http.StatusOK {
		c <- response{data: nil, error: fmt.Errorf("Skipping Commic : %d, Got: %d\n", counter, res.StatusCode)}
		return
	}
	
	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		c <- response{data: nil, error: fmt.Errorf("Invalid Body: %d\n", err)}
		return
	}
	
	res.Body.Close()
	
	c <- response{data: body, error: nil}
}

func write(f *os.File, data chan response, counter int) {
	response := <-data
	
	if response.error != nil {
		fmt.Printf("failed: #%d\n", counter)
		return
	}
	
	if counter > 1 {
		f.WriteString(",\n")
	}
	
	f.Write(response.data)
}

func GetData(iteration int) {
	start := time.Now()
	
	data := make(chan response, 32)
	wg := sync.WaitGroup{}
	
	f, err := os.Create("./outputs/concurrent.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't Create the file")
	}
	
	defer f.Close()
	
	f.WriteString("[\n")
	
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		go func(counter int) {
			defer wg.Done()
			getOne(data, counter)
			write(f, data, counter)
		}(i)
	}
	
	wg.Wait()
	
	f.WriteString("\n]")
	
	secs := time.Since(start).Seconds()
	fmt.Printf("It tool: %.2v seconds | %.2v mins\n", secs, secs/60)
}
