package main

import (
	"fmt"
	"sync"
)

type UrlCounter struct {
	mu sync.Mutex
	wg sync.WaitGroup
	v  map[string]int
}

func (u *UrlCounter) Inc(key string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.v[key]++
}

func (u *UrlCounter) Exist(key string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.v[key] == 1
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(counter *UrlCounter, url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 || counter.Exist(url) {
		return
	}
	counter.Inc(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		counter.wg.Add(1)
		go func() {
			Crawl(counter, u, depth-1, fetcher)
			defer counter.wg.Done()
		}()
	}
	return
}

func main() {
	counter := &UrlCounter{v: make(map[string]int)}
	Crawl(counter, "https://golang.org/", 4, fetcher)
	counter.wg.Wait()
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
