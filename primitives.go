package main

func Button(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "button",
		text: text,
	}
}

func P(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "p",
		text: text,
	}
}

func H1(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h1",
		text: text,
	}
}

func H2(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h2",
		text: text,
	}
}

func H3(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h3",
		text: text,
	}
}

func H4(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h4",
		text: text,
	}
}

func H5(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h5",
		text: text,
	}
}

func H6(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "h6",
		text: text,
	}
}

func Div(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "div",
		children: children,
	}
}

func Body(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "body",
		children: children,
	}
}

func Nav(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "nav",
		children: children,
	}
}

func Aside(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "aside",
		children: children,
	}
}

func Section(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "section",
		children: children,
	}
}

func Main(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "main",
		children: children,
	}
}

func Link(href string, rel string) *HtmlElement {
	return &HtmlElement{
		tag: "link",
		attr: map[string]string{
			"href": href,
			"rel":  rel,
		},
	}
}

func Head(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "head",
		children: children,
	}
}

func A(link string, text string, children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "a",
		text:     text,
		children: children,
		attr: map[string]string{
			"href": link,
		},
	}
}

func Row(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "div",
		children: children,
		attr: map[string]string{
			"class": "row spacing-md",
		},
	}
}
