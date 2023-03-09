package Builder

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

type Person struct {
	StreetAddress, Postcode, City string

	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (p *PersonBuilder) Builder() *Person {
	return p.person
}

func (p *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	p.person.StreetAddress = street
	return p
}
func (p *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	p.person.City = city
	return p
}
func (p *PersonAddressBuilder) PostCode(code string) *PersonAddressBuilder {
	p.person.Postcode = code
	return p
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (p *PersonJobBuilder) At(comapny string) *PersonJobBuilder {
	p.person.CompanyName = comapny
	return p
}

func (p *PersonJobBuilder) Asa(job string) *PersonJobBuilder {
	p.person.Position = job
	return p
}

func (p *PersonJobBuilder) Earn(income int) *PersonJobBuilder {
	p.person.AnnualIncome = income
	return p
}

func Start_BUilder() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London").
		In("London").
		PostCode("Mortgatan6").
		Works().
		At("Plea").
		Asa("Programmer").
		Earn(12300)

	person := pb.Builder()
	fmt.Println(person)
}

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	e.email.from = from
	return e
}

func (e *EmailBuilder) To(to string) *EmailBuilder {
	e.email.to = to
	return e
}

func (e *EmailBuilder) Subject(to string) *EmailBuilder {
	e.email.subject = to
	return e
}

func (e *EmailBuilder) Body(to string) *EmailBuilder {
	e.email.body = to
	return e
}

func sendMailImpl(email *email) {
	SendEmail(func(b *EmailBuilder) {
		b.From("guiowoo").To("holymoly").
			Body("hello there").Subject("Hihi")
	})
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

type Student struct {
	name, position string
}
type StdMod func(*Student)
type StdBuilder struct {
	actions []StdMod
}

func (b *StdBuilder) Called(name string) *StdBuilder {
	b.actions = append(b.actions, func(student *Student) {
		student.name = name
	})
	return b
}

func (b *StdBuilder) SetPos(pos string) *StdBuilder {
	b.actions = append(b.actions, func(student *Student) {
		student.position = pos
	})
	return b
}

func (b *StdBuilder) Build() *Student {
	p := Student{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func Start2() {
	b := StdBuilder{}
	p := b.Called("Guiwoo").SetPos("Lv3").Build()
	fmt.Println(p)
}
