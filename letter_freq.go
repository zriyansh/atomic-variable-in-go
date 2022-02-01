package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()              // to close the get request
	body, _ := ioutil.ReadAll(resp.Body) // to read body part of url we fetch
	for _, b := range body {             // reads every byte of the body
		c := strings.ToLower(string(b))
		// we are only dealing with lowercase, and as 'b' is in byte, we convert it to string
		index := strings.Index(allLetters, c) // returns index (in allLetters) of c.
		if index >= 0 {                       // to check if char is b/w a-z only, not special chars.
			frequency[index] += 1
		}
	}
}

func main() {
	var frequency [26]int32
	for i := 1000; i < 1003; i++ {
		// 1000 so that our URL address can be matched
		countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)

	}
	println("done")
	for i, f := range frequency {
		fmt.Printf("%s %d\n", string(allLetters[i]), f)
	}
}
