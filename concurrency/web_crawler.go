package concurrency

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var errLoading = errors.New("url is loading")

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Println("depth limit reached")
		return
	}

	// 检查是否已经抓取过或正在抓取
	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		// URL 已经被抓取过或正在抓取，直接返回
		fetched.Unlock()
		return
	}
	// 标记为正在抓取
	fetched.m[url] = errLoading
	fetched.Unlock()

	body, urls, err := fetcher.Fetch(url)

	// 更新抓取状态
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// 使用 WaitGroup 来等待所有 goroutines 完成
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			Crawl(url, depth-1, fetcher)
		}(u)
	}
	wg.Wait()
}

func RunWebCrawler() {
	Crawl("http://golang.org/", 4, fetcher)
	fmt.Println("Fetching stats\n--------------")
	for url, err := range fetched.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
