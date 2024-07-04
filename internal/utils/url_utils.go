package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func FetchContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RemoveDuplicatesAndNormalize(urls []string, baseURL string) []string {
	seen := make(map[string]bool)
	var result []string

	baseURLParsed, err := url.Parse(baseURL)
	if err != nil {
		return urls // Return original URLs if base URL parsing fails
	}

	for _, u := range urls {
		normalized, err := NormalizeURL(u, baseURLParsed)
		if err != nil {
			continue // Skip invalid URLs
		}

		if !seen[normalized] {
			seen[normalized] = true
			result = append(result, normalized)
		}
	}

	return result
}

func NormalizeURL(u string, baseURL *url.URL) (string, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	// Resolve relative URLs
	resolvedURL := baseURL.ResolveReference(parsedURL)

	// Normalize the URL
	resolvedURL.RawQuery = "" // Remove query parameters
	resolvedURL.Fragment = "" // Remove fragments

	// Ensure the scheme is set
	if resolvedURL.Scheme == "" {
		resolvedURL.Scheme = baseURL.Scheme
	}

	return resolvedURL.String(), nil
}
