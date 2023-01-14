package normal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getOne(counter int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", counter)
	res, err := http.Get(url)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant get to the webpage\n")
		os.Exit(-1)
	}
	
	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Skipping Commic : %d, Got: %d\n", counter, res.StatusCode)
		return nil
	}
	
	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Body: %d\n", err)
		os.Exit(-1)
	}
	
	res.Body.Close()
	
	return body
}

func GetData(iteration int) {
	start := time.Now()
	var fails, i int
	
	f, err := os.Create("./outputs/normal.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't Create the file")
	}
	
	f.WriteString("[\n")
	
	for i = 1; i <= iteration; i++ {
		if fails > 1 {
			break
		}
		
		data := getOne(i)
		
		if data == nil {
			fails++
			continue
		}
		
		if i > 1 {
			f.WriteString(",\n")
		}
		
		f.Write(data)
		fails = 0
	}
	
	f.WriteString("\n]")
	f.Close()
	
	secs := time.Since(start).Seconds()
	fmt.Printf("Read Total Number Of %d Comics\n", i)
	fmt.Printf("It tool: %v seconds | %v mins\n", secs, secs/60)
}
