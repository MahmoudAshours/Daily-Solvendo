package main

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("!(&(f,t))|(f,t)"))
}
func parseBoolExpr(expression string) bool {
	stk := []rune{} 

	for _, c := range expression {
		if c != ')' && c != ',' {
			stk = append(stk, c)
		} else if c == ')' { // When ')' is encountered, evaluate subexpression
			exp := []bool{} // Slice to hold boolean values of the current subexpression

			// Pop characters until '(' is found, collect 't' or 'f' values
			for len(stk) > 0 && stk[len(stk)-1] != '(' {
				t := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				exp = append(exp, t == 't')
			}

			stk = stk[:len(stk)-1] // Pop the '(' from the stack

			if len(stk) > 0 {
				t := stk[len(stk)-1] // Get the operator before '('
				stk = stk[:len(stk)-1]
				v := exp[0] // Initialize result with the first value

				// Perform the corresponding logical operation
				if t == '&' { // AND operation
					for _, b := range exp {
						v = v && b
					}
				} else if t == '|' { // OR operation
					for _, b := range exp {
						v = v || b
					}
				} else { // NOT operation
					v = !v
				}

				if v {
					stk = append(stk, 't')
				} else {
					stk = append(stk, 'f')
				}
			}
		}
	}

	// Return the final result from the stack
	return stk[len(stk)-1] == 't'
}
