package main

import "fmt"

// OCP
// open for Extension, closed for modification
// Specification
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (k SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == k.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(product []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range product {
		if spec.IsSatisfied(&v) {
			result = append(result, &product[i])
		}
	}
	return result
}

var _ Specification = (*SizeSpecification)(nil)
var _ Specification = (*ColorSpecification)(nil)

func main() {
	apple := Product{"apple", green, small}
	tree := Product{"tree", green, large}
	house := Product{"hosue", blue, large}

	producs := []Product{apple, tree, house}

	f := &Filter{}
	for _, v := range f.FilterByColor(producs, green) {
		fmt.Printf("- %s is green \n", v.name)
	}
}
