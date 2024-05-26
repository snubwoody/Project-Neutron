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

func button(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "button",
		text: text,
	}
}

func p(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "p",
		text: text,
	}
}

func h1(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h1",
		text: text,
	}
}

func h2(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h2",
		text: text,
	}
}

func h3(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h3",
		text: text,
	}
}

func h4(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h4",
		text: text,
	}
}

func h5(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h5",
		text: text,
	}
}

func h6(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h6",
		text: text,
	}
}

func div(children []*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "div",
		children: children,
	}
}
