package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type urlFetched struct {
	urls map[string]bool
	lock sync.Mutex
}

func (u *urlFetched) isFetched(url string) bool {
	u.lock.Lock()
	defer u.lock.Unlock()
	if _, ok := u.urls[url]; ok {
		return true
	}
	u.urls[url] = true
	return false
}

var lock_map *urlFetched = &urlFetched{urls: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, parent_chan chan bool) {
	defer func() {
		// send to channel
		parent_chan <- true
	}()
	if depth <= 0 {
		return
	}
	if lock_map.isFetched(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	child_chan := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, child_chan)
	}
	for range urls {
		// receive from channel
		<-child_chan
	}
	return
}

func main() {
	ch := make(chan bool)
	// Exactly: since no go routine is waiting to receive, the send is blocked, and your program is deadlocked.
	// go Crawl("https://golang.org/", 4, fetcher, ch)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
