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
			Aside(
				P("Hello world"),
				Button("Click me"),
			),
			Section(
				P("Hello world"),
				Button("Click me"),
				A("https://youtube.com", "This is a link to youtube"),
				Row(
					P("Hi"),
					P("Hello"),
					P("World"),
				),
			),
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
