package main

import "fmt"

func main() {
	htmlString := "<html><body><h1>Title</h1><p>Paragraph</p></body></html>"

	invalidHTML := "<html><body><h1>Title<p>Paragraph</h1></body></html>"
	fmt.Println(validateHTML(htmlString))
	fmt.Println(validateHTML(invalidHTML))
}

func validateHTML(html string) bool {
	var stack Stack
	for i, v := range html {
		if v == '<' {
			if html[i+1] == '/' {
				stack.Pop()
			} else {
				stack.Push("<")
			}
		}
	}
	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return "empty"
	}
	// Get the last element
	item := s.items[len(s.items)-1]
	// Remove the last element
	s.items = s.items[:len(s.items)-1]
	return item
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
