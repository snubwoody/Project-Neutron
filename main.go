package main

import (
	"bufio"
	"os"
)

func main() {
	Page("index.html")
}

func Page(path string) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		panic("Couldn't open html file")
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	w.Write([]byte("<!DOCTYPE html>\n"))

	h1 := &HtmlElement{
		tag:  "h1",
		text: "This is a heading",
	}

	p := &HtmlElement{
		tag:  "p",
		text: "This is a paragraph",
	}

	div := &HtmlElement{
		tag:      "div",
		class:    []string{"div"},
		children: []*HtmlElement{h1, p},
	}

	head := &HtmlElement{
		tag: "head",
	}

	body := &HtmlElement{
		tag:      "body",
		class:    []string{"hi"},
		children: []*HtmlElement{div},
	}

	html := HtmlElement{
		tag:      "html",
		children: []*HtmlElement{head, body},
	}

	html.render(w)
}
