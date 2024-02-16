package main

import (
	"fmt"
	"sync"
)

var visitedURLs map[string]bool = make(map[string]bool)
var visitedURLsMtx sync.Mutex = sync.Mutex{}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, parentThreadWaitGroup *sync.WaitGroup) {
	var crawlSubThreadsWaitGroup sync.WaitGroup = sync.WaitGroup{}

	if parentThreadWaitGroup != nil {
		defer parentThreadWaitGroup.Done()
	}

	if depth <= 0 {
		return
	}

	_, ok := visitedURLs[url]

	if !ok {
		visitedURLs[url] = true
		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			visitedURLs[url] = false
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			crawlSubThreadsWaitGroup.Add(1)
			go Crawl(u, depth-1, fetcher, &crawlSubThreadsWaitGroup)
		}

		crawlSubThreadsWaitGroup.Wait()
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher, nil)
}

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
