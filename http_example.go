package main

import (
	"fmt"
	"net/http"
)

func getHttp(url string) (string, error) {

	resp, err := http.Get(url)

	if err != nil {

		return "", err
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {

		return "", fmt.Errorf("Content Type is nil")
	}

	return ctype, nil

}

func main() {

	c, e := getHttp("https://google.com")
	if e == nil {

		fmt.Println(c)
	} else {

		fmt.Println(e)
	}
}
