package stack

import "fmt"

func ValidParentheses() {
	variant1 := "(){}[]"
	variant2 := "(]"
	fmt.Printf("Is valid this string %s: %t\n", variant1, isValid(variant1))
	fmt.Printf("Is valid this string %s: %t\n", variant2, isValid(variant2))
}

func isValid(s string) bool {
	// first we define stack data type by using go's slices
	var stack Stack

	// then we need a HaspMap for coupling open and closed brackets
	closeToOpenHashMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		// here we check if current value is closing parentheses, because every key in our hasp map is a closing rune
		if _, ok := closeToOpenHashMap[c]; ok {
			// check that stack is not empty and last rune equal to closing parentheses
			if !stack.isEmpty() && stack[len(stack)-1] == closeToOpenHashMap[c] {
				stack.Pop()
			} else {
				return false
			}
		} else {
			stack.Push(c)
		}
	}

	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

type Stack []rune

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str rune) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (rune, bool) {
	if s.isEmpty() {
		return -1, false
	} else {
		topElemIdx := len(*s) - 1
		topElem := (*s)[topElemIdx]
		*s = (*s)[:topElemIdx]
		return topElem, true
	}
}
