package main

import (
	"bufio"
	"os"
)

func main() {
	Page("dist/index.html")
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

	body := Body(
		Nav(
			P("Dashboard"),
		),
		Main(
			Sidebar(),
			Section(),
		),
	)

	head := Head(
		Link("styles/preflight.css", "stylesheet"),
		Link("styles/base.css", "stylesheet"),
	)

	html := HtmlElement{
		tag:      "html",
		children: []*HtmlElement{head, body},
	}

	html.render(w)
}

func Sidebar() *HtmlElement {
	return Column(
		Row(
			P("Top repositories"),
			Button("New"),
		),
		A("#", "snubwoody/repo"),
		A("#", "snubwoody/repo"),
		A("#", "snubwoody/repo"),
		A("#", "snubwoody/repo"),
		A("#", "snubwoody/repo"),
		A("#", "snubwoody/repo"),
	)
}
