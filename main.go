package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	client := &http.Client{}
	// c := make(chan string, len(os.Args)-1)
	var wg sync.WaitGroup
	// go receive(client, c, &wg)

	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go func(symbol string, client *http.Client) {
			defer wg.Done()
			getNextEPS(symbol, client)
		}(arg, client)
	}

	wg.Wait()
}

// func receive(client *http.Client, c chan string, wg *sync.WaitGroup) {
// 	for symbol := range c {
// 		wg.Add(1)
// 		go func() {
// 			getNextEPS(symbol, client)
// 			wg.Done()
// 		}()
// 	}
// }

// ingoring all the errors because I am cool like that!
func getNextEPS(symbol string, client *http.Client) {
	// preparing request
	pathElems := []string{"https://www.earningswhispers.com/api/getstocksdata/", symbol}
	path := strings.Join(pathElems, "")
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", "https://www.earningswhispers.com/stocks/STEM")

	// performing http req
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	// reading body as bytes
	body, _ := io.ReadAll(resp.Body)

	// converting resp in GENERIC json
	var j interface{}
	json.Unmarshal(body, &j)

	// converting json to map
	m := j.(map[string]interface{})

	// printing needed value from map
	mess := []string{symbol, "\t --> \t", m["nextEPSDate"].(string), "\n"}
	fmt.Printf(strings.Join(mess, ""))
}
