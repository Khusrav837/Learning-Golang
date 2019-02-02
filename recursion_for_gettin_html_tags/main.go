package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	url = strings.Trim(url, "\r\n")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get Url err %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	return links
}
