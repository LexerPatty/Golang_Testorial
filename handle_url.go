/*URL handling is very common in Go, especially for APIs, backend, and web apps.

Scheme: The protocol used by the URL (http, https)
Host: The domain name of the URL
Path: The resource path
RawQuery: The query parameters present after the ? symbol
*/

package main

import (
	"fmt"
	"net/url"
)

func main() {
    rawURL := "https://example.com/search?q=golang&page=2"

    u, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Invalid URL")
        return
    }

    fmt.Println("Scheme:", u.Scheme)
    fmt.Println("Host:", u.Host)
    fmt.Println("Path:", u.Path)
    fmt.Println("Query:", u.RawQuery)
	// This code parses a URL and extracts its components like scheme, host, path, and query parameters.

	
	// Building a URL...................................................................................
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/search",       

		RawQuery: url.Values{
			"q":   {"golang"},
			"page": {"2"},
		}.Encode(),

		// q := baseURL.Query()
		// q.Set("q", "golang")
		// q.Set("page", "1")
	
		// baseURL.RawQuery = q.Encode()
	}


	fmt.Println("Built URL:", baseURL.String())


	// 2. Reading Query Parameters:.........................................................................
	values, err := url.ParseQuery(u.RawQuery)
	fmt.Println(values.Get("q"))     // golang
	fmt.Println(values.Get("page"))  // 2


	// MODIFYING URL COMPONENTS:...................................................................................
	u.Path = "/new-search"
	u.RawQuery = "pushpendu-dey"

	fmt.Println("Modified URL:", u.String())
}