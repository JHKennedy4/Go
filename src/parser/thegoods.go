package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	// client := new(http.Client)
	resp, err := http.Get("http://www.theatlantic.com/international/archive/2014/12/do-police-body-cameras-work-ferguson/383323/")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := html.Parse(resp.Body)
	// walk(doc)
	doc, _ = nodeRunner(doc)
	printTree(doc)
}

func printTree(n *html.Node) {
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		fmt.Println(strings.TrimSpace(n.Data))
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTree(c)
	}
	return
}

func nodeRunner(n *html.Node) (*html.Node, int) {

	switch n.Type {
	case html.ElementNode:
		// at this level, we determine wether to inlude this node by examining it's children
		var pre_c html.Node
		var a_length, c_length int

		c := n.FirstChild

		for c != nil {
			var length int
			// do the thing where we decide what type it is
			run_c, length := nodeRunner(c)
			if run_c != nil {
				c = run_c
			} else {
				if c == n.FirstChild {
					n.FirstChild = c.NextSibling
				} else {
					pre_c.NextSibling = c.NextSibling
				}
			}

			if length != -1 {
				if run_c.Data == "a" {
					a_length += length
				} else {
					c_length += length
				}
			}
			// this part deletes a child node
			if true {
				pre_c.NextSibling = c.NextSibling
			}

			pre_c = *c
			c = c.NextSibling
		}
		return n, -1
	case html.TextNode:
		// this is a terminating case.
		// the data I need is the length of the text
		//
		// an element node block that is parsing it's children recieves that length
		return n, len(strings.TrimSpace(n.Data))
	}

	return n, -1
}

/*
Terminating case of recursion is always a text node
Delete parent when sibling anchor text node length is greater than sibling text node length

Need to reassemble the HTML tree minus the eliminated nodes
*/

type NodeTextLength struct {
	html.Node
	TextLength int
}

func ProcessElement(n *html.Node) int {
	// current node n.Data
	var textLength int
	// pring open of current node
	fmt.Printf("<%s>\n", n.Data)
	// iterate across child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch {
		case n.Data == "a":
			//linkLength += walk(c)
		default:
			textLength += walk(c)
		}
	}
	fmt.Sprintf("text length: %d", textLength)
	// pring close of current node
	//fmt.Printf("</%s>", n.Data)
	return textLength

}

func ProcessText(n *html.Node) int {
	// fmt.Println(string(len(n.Data)))
	data := strings.TrimSpace(n.Data)
	if len(data) > 0 {
		fmt.Println(n.Data)
		fmt.Printf("Length: %d\n\n", len(data))
	}
	return len(data)
}

// terminal case is TextNode
func walk(n *html.Node) int {
	switch n.Type {
	case html.ElementNode:
		return ProcessElement(n)
	case html.TextNode:
		return ProcessText(n)
	case html.ErrorNode:
	default:
		return ProcessElement(n)
	}

	return 0
}
