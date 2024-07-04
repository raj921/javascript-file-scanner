package scanner

import (
	"regexp"
	"sync"

	"github.com/raj921/javascript-file-scanner/internal/utils"
)

var (
	urlPatterns = map[string]*regexp.Regexp{
		"URL":        regexp.MustCompile(`(?i)(?:url|src|href)\s*[=:]\s*["']?([^"'\s>]+)`),
		"JavaScript": regexp.MustCompile(`(?i)["']([^"']+\.js)`),
		"CSS":        regexp.MustCompile(`(?i)["']([^"']+\.css)`),
		"Image":      regexp.MustCompile(`(?i)["']([^"']+\.(png|jpg|jpeg|gif|svg|webp))`),
		"API":        regexp.MustCompile(`(?i)["']([^"']+/api/[^"']+)`),
		"DataURI":    regexp.MustCompile(`data:[^;,\s]+[;,]`),
	}
)

func ExtractURLs(content, baseURL string) map[string][]string {
	var wg sync.WaitGroup
	results := make(map[string][]string)
	resultChan := make(chan struct {
		category string
		urls     []string
	}, len(urlPatterns))

	for category, pattern := range urlPatterns {
		wg.Add(1)
		go func(cat string, pat *regexp.Regexp) {
			defer wg.Done()
			matches := pat.FindAllStringSubmatch(content, -1)
			var urls []string
			for _, match := range matches {
				if len(match) > 1 {
					urls = append(urls, match[1])
				}
			}
			resultChan <- struct {
				category string
				urls     []string
			}{cat, urls}
		}(category, pattern)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.category] = utils.RemoveDuplicatesAndNormalize(result.urls, baseURL)
	}

	return results
}
