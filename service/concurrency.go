package service

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func ConcurrentCheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, checker(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

func LoopCheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = checker(url)
	}

	return results
}
