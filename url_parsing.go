package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"
	parseURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return
	}
	fmt.Println("Scheme:", parseURL.Scheme)
	fmt.Println("Host:", parseURL.Host)
	fmt.Println("Port:", parseURL.Port())
	fmt.Println("Path:", parseURL.Path)
	fmt.Println("RawQuery:", parseURL.RawQuery)
	fmt.Println("Fragment:", parseURL.Fragment)

	rawURL1 := "https://example.com/path?name=Simon&age=13"

	parseURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return
	}

	queryParam := parseURL1.Query()
	fmt.Println(queryParam)
	fmt.Println("Name:", queryParam.Get("name"))
	fmt.Println("Age:", queryParam.Get("age"))

	// Building a URL

	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	query := baseURL.Query()
	query.Set("name", "Simon")
	query.Set("age", "13")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Build URL:", baseURL.String())

	values := url.Values{}
	// Add key-values pairs to the values obkect
	values.Add("name", "Simon")
	values.Add("age", "13")
	values.Add("city", "Santiago")
	values.Add("country", "Chile")

	// Encode
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	// Build a URL
	baseURL1 := "http://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery
	fmt.Println("URL:", fullURL)

}
