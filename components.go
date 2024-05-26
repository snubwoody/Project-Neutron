package main

func Row(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "div",
		children: children,
		attr: map[string]string{
			"class": "row spacing-md",
		},
	}
}

func Column(children ...*HtmlElement) *HtmlElement {
	return &HtmlElement{
		tag:      "div",
		children: children,
		attr: map[string]string{
			"class": "column spacing-md",
		},
	}
}

func Button(text string) *HtmlElement {
	return &HtmlElement{
		tag:  "button",
		text: text,
		attr: map[string]string{
			"class": "button-primary",
		},
	}
}
