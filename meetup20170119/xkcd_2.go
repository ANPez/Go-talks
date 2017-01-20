// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://xkcd.com/%d/info.0.json"

func getURL(id int) (string, error) {
	u := fmt.Sprintf(url, id)
	resp, err := http.Get(u)
	if nil != err {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(body, &obj)
	if nil != err {
		return "", err
	}

	return obj["img"].(string), nil
}

// INIT START OMIT

func main() {
	requests := make(chan int)
	exitCh := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func() {
			for id := range requests {
				url, err := getURL(id)
				if nil != err {
					log.Println(err)
				} else {
					log.Printf("Got URL for id %d: %s", id, url)
				}
			}
			exitCh <- struct{}{}
		}()
	}
	// INIT END OMIT

	// Process START OMIT

	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests) // HL

	for i := 0; i < 3; i++ {
		<-exitCh
	}
	// Process END OMIT
}
