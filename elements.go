package main

import (
	"bufio"
	"fmt"
	"html/template"
	"strings"
)

type HtmlElement struct {
	tag      string
	text     string
	class    []string
	children []*HtmlElement
}

func (element *HtmlElement) render(w *bufio.Writer) {
	//FIXME class names on null classes
	htmlElement, _ := template.New(element.tag).Parse(
		fmt.Sprintf(
			"<%s class={{.class}}>{{.text}}\n",
			element.tag,
		),
	)

	templateAttr := map[string]string{
		"class": fmt.Sprintf("%s", strings.Join(element.class, " ")),
		"text":  element.text,
	}

	if writeErr := htmlElement.Execute(w, templateAttr); writeErr != nil {
		panic(writeErr)
	}

	if element.children != nil {
		for _, element := range element.children {
			element.render(w)
		}
	}

	w.Write([]byte(fmt.Sprintf("</%s>\n", element.tag)))

	w.Flush()
}

func (element *HtmlElement) appendChildren(elements []*HtmlElement) {
	element.appendChildren(elements)
}
