package designPatern

import (
	"fmt"
	"strings"
)

// flexible structure sort of object instead steps

const indentSize = 2

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName, HtmlElement{rootName, "", []HtmlElement{}}}
}

func StartTest() {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}
