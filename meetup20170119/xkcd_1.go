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

// getURL END OMIT

func main() {
	for i := 1; i <= 10; i++ {
		url, err := getURL(i)
		if nil != err {
			log.Println(err)
		} else {
			log.Printf("Found url for XKCD %d: %s", i, url)
		}
	}
}
