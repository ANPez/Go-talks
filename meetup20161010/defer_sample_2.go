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
	defer resp.Body.Close() // HL

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if nil != err {
		return "", err
	}

	return result["ip"], nil
}

// getIPAddress END OMIT

func main() {
	fmt.Println(getIPAddress())
}
