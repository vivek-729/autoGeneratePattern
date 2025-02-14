package main

import (
	"github.com/vivek-729/autoGeneratePattern/code"
	"fmt"
)

// type stack struct {
// 	arr []byte
// }

// func (s *stack) empty() bool {
// 	return len(s.arr) == 0
// }

// func (s *stack) push(a byte) {
// 	s.arr = append(s.arr, a)
// }

// func (s *stack) pop() bool {
// 	if len(s.arr) > 0 {
// 		s.arr = s.arr[:len(s.arr)-1]
// 		return true
// 	}
// 	return false
// }

// func (s *stack) top() byte {
// 	return s.arr[len(s.arr)-1]
// }

func main() {
	code.AutoGenerate()
	// infix to postfix dummy coDe
	// s := "((A+B)-C*(D/E))+F"
	fmt.Println("Hello! world")

	// operator := []byte{'-', '+', '*', '/'}
	// precedence := make(map[byte]float32)
	// precedence['+'] = 11.1
	// precedence['-'] = 11.2
	// precedence['*'] = 12.1
	// precedence['/'] = 12.2

	// var st stack
	// var ans string

	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		st.push('(')
	// 	} else if s[i] == ')' {
	// 		for !st.empty() && st.top() != '(' {
	// 			ans += string(st.top())
	// 			st.pop()
	// 		}
	// 		st.pop()
	// 	} else if slices.Contains(operator, s[i]) {
	// 		for !st.empty() && precedence[s[i]] <= precedence[st.top()] {
	// 			ans += string(st.top())
	// 			st.pop()
	// 		}
	// 		st.push(s[i])
	// 	} else {
	// 		ans += string(s[i])
	// 	}
	// }

	// for !st.empty() {
	// 	ans += string(st.top())
	// 	st.pop()
	// }

	// fmt.Println(ans)
}
