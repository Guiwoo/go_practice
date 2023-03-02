package main

import (
	"fmt"
	"os"
)

type Student struct {
	name string
}

func (s *Student) String() string {
	return fmt.Sprintf("One for %s,one for me", s.name)
}

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	s := &Student{"you"}
	if len(name) > 0 {
		s.name = name
	}
	return s.String()
}

type Testing interface {
	nothing()
}
type Abc struct {
	a Testing
}

func (a *Abc) nothing() {}

type B struct{}

func (b *B) nothing() {}

func hello(a Testing) *B {
	res := a.(*B)
	return res
}

func main() {
	os.WriteFile("filename", []byte("Something happens on the internet"), 0644)
}
