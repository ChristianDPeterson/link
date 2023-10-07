package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	htmlFileName := "ex2.html"
	htmlFile, err := os.Open(htmlFileName)
	if err != nil {
		log.Fatal(err)
	}

	document, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatal(err)
	}

	dfs(document)

}

func dfs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c)
	}
}
