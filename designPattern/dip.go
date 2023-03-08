package designPatern

import "fmt"

// HLM should not depend on LLM

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}

	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level-module

type Research struct {
	//break Dip
	relationships Relationships
	browser       RelationshipBrowser
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

var _ RelationshipBrowser = (*Relationships)(nil)

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
