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
	attr     map[string]string
	class    []string
	children []*HtmlElement
}

func (element *HtmlElement) render(w *bufio.Writer) {
	attr := parseAttr(element.attr)
	htmlElement, _ := template.New(element.tag).Parse(
		fmt.Sprintf(
			"<%s %s>{{.text}}\n",
			element.tag,
			attr,
		),
	)

	templateAttr := map[string]string{
		"text": element.text,
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

func parseAttr(attr map[string]string) string {
	var attributes []string
	for k, v := range attr {
		attributes = append(attributes, fmt.Sprintf("%s='%s'", k, v))
	}

	return strings.Join(attributes, " ")
}
