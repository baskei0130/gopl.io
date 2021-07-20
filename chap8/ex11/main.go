package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetch(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	if cancelled() {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("failed to create request to %s: %v\n", url, err)
		return
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to get response from %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read %s: %v\n", url, err)
		return
	}

	if !cancelled() {
		ch <- string(b)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch, &wg)
	}

	fmt.Println(<-ch)
	close(done)

	wg.Wait()
	close(ch)
}
