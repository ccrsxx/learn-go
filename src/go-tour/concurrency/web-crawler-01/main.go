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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, cachedUrl *CachedUrl) {
	defer wg.Done()

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if cachedUrl.Visit(url) {
		fmt.Printf("already visited: %s\n", url)
		return
	}

	fmt.Printf("visiting: %s\n", url)

	body, urls, err := fetcher.Fetch(url)

	fmt.Printf("Crawl: depth=%d url=%s\n", depth, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, url := range urls {
		wg.Add(1)
		go Crawl(url, depth-1, fetcher, wg, cachedUrl)
	}
}

type CachedUrl struct {
	mu          sync.Mutex
	visitedUrls map[string]bool
}

func (c *CachedUrl) Visit(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.visitedUrls[url] {
		return true // Already visited
	}

	c.visitedUrls[url] = true // Mark it NOW, while we still hold the lock

	return false // New visit
}

func main() {
	var wg sync.WaitGroup

	cachedUrl := CachedUrl{visitedUrls: make(map[string]bool)}

	wg.Add(1)

	go Crawl("https://golang.org/", 4, fetcher, &wg, &cachedUrl)

	wg.Wait()
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
