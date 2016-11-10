// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIPAddress() (string, error) {
	resp, err := http.Get("https://api.ipify.org/?format=json")
	if nil != err {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		resp.Body.Close() // HL
		return "", err
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if nil != err {
		resp.Body.Close() // HL
		return "", err
	}

	resp.Body.Close() // HL
	return result["ip"], nil
}

// getIPAddress END OMIT

func main() {
	fmt.Println(getIPAddress())
}
