// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "http://xkcd.com/%d/info.0.json"

func getURL(comicID int) (string, error) {
	url := fmt.Sprintf(URL, comicID)
	resp, err := http.Get(url)
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
type response struct {
	ID  int
	URL string
	Err error
}

func main() {
	c1 := make(chan int)
	c2 := make(chan response)
	cExit := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func() { // HL
			for id := range c1 {
				url, err := getURL(id)
				c2 <- response{id, url, err}
			}
			cExit <- struct{}{}
		}()
	}
	// INIT END OMIT

	// Process START OMIT
	go func() {
		for r := range c2 {
			if nil != r.Err {
				log.Println(r.Err)
			} else {
				log.Printf("Found url for XKCD %d: %s", r.ID, r.URL)
			}
		}
		cExit <- struct{}{}
	}()

	for i := 1; i <= 10; i++ {
		c1 <- i
	}
	close(c1) // HL

	for i := 0; i < 3; i++ {
		<-cExit
	}

	close(c2) // HL
	<-cExit
	// Process END OMIT
}
