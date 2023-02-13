package main

import (
	chapter27 "awesomeProject/part_2/chapter27"
	"fmt"
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

func main() {
	a := &chapter27.Email{}
	//b := &chapter27.FinanceReport{"abcd"}
	var c *chapter27.Report
	chapter27.SendReport(c, a, "park.guiwoo@hotmail.com")
}
