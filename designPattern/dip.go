package designPatern

import "fmt"

// HLM should not depend on LLM

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Personer struct {
	name string
}

type Info struct {
	from         *Personer
	relationship Relationship
	to           *Personer
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Personer {
	result := make([]*Personer, 0)

	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}

	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Personer) {
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
	FindAllChildrenOf(name string) []*Personer
}

var _ RelationshipBrowser = (*Relationships)(nil)

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
