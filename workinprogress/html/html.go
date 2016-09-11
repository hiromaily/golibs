package html

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

//TODO:work in progress

func checkElement(t html.Token, tag string) {
	//fmt.Println(t.Data)
	isMatch := t.Data == tag
	if isMatch {
		switch tag {
		case "div":

		case "a":
			checkAttr(t, "href", "http://textream.yahoo.co.jp/rd/finance/7893")
		case "input":
			fmt.Println("input")
			checkAttr(t, "name", "gintoken")
			checkAttr(t, "name", "fr")
		default:
			//fmt.Println("default")
		}

	}
}

func checkAttr(t html.Token, attr, val string) {
	// Get Tag Attribute
	for _, a := range t.Attr {
		if a.Key == attr {
			if a.Val == val {
				fmt.Printf("Found %s:%s\n", attr, a.Val)
			}
			break
		}
	}
}

// ParseHTTPBody is to parse HTTP request body
func ParseHTTPBody(body io.ReadCloser) *html.Tokenizer {
	z := html.NewTokenizer(body)
	return z
}

// ParseHTMLText is to parse HTML text data
func ParseHTMLText(text string) (*html.Node, error) {
	//*Node
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// ParseHTMLText2 is to parse HTML text data using html.Node
func ParseHTMLText2(text string) ([]*html.Node, error) {
	var context html.Node
	context = html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	}
	//[]*Node
	nodes, err := html.ParseFragment(strings.NewReader(text), &context)

	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// ParseToken is to parse token
func ParseToken(z *html.Tokenizer, tag string) {
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// check element
			checkElement(t, tag)
		}
	}
}

// ParseNode is to parse node
// TODO: work in progress
func ParseNode(z *html.Node, tag string) {

}
