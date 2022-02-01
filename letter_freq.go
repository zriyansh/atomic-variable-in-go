package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

// var lock = sync.Mutex{}

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()              // to close the get request
	body, _ := ioutil.ReadAll(resp.Body) // to read body part of url we fetch
	for _, b := range body {             // reads every byte of the body
		c := strings.ToLower(string(b))
		// we are only dealing with lowercase, and as 'b' is in byte, we convert it to string
		// lock.Lock()
		index := strings.Index(allLetters, c) // returns index (in allLetters) of c.
		if index >= 0 {                       // to check if char is b/w a-z only, not special chars.
			atomic.AddInt32(&frequency[index], 1)
		}
		// lock.Unlock()
	}
	wg.Done()
}

// atomic variable implemenation took 3 sec
// mutex(multithread) implementation took 27 sec
// single threaded implemenation took 37 seconds
// we used atomic var bcoz we were not context switching many times over, else this would have not been a good option
// in this program, our countLetters () is fast, we have to just wait for pages to load that is taking all the time

func main() {
	var frequency [26]int32
	wg := sync.WaitGroup{}
	for i := 1000; i < 1200; i++ {
		wg.Add(1)
		// 1000 so that our URL address can be matched
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &wg)

	}
	wg.Wait()
	println("done")
	for i, f := range frequency {
		fmt.Printf("%s %d\n", string(allLetters[i]), f)
	}
}
