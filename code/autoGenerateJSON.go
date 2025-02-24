package code

import (
	"encoding/json"
	"github.com/vivek-729/autoGeneratePattern/node"
	"fmt"
	"os"
	"slices"
)

type stack struct {
	arr []string
}

func (s *stack) empty() bool {
	return len(s.arr) == 0
}

func (s *stack) push(a string) {
	s.arr = append(s.arr, a)
}

func (s *stack) pop() bool {
	if len(s.arr) > 0 {
		s.arr = s.arr[:len(s.arr)-1]
		return true
	}
	return false
}

func (s *stack) top() string {
	return s.arr[len(s.arr)-1]
}

func AutoGenerate() {
	// fmt.Println("Hi")
	// fmt.Printf("%+v\n", variable.V20)
	// fmt.Printf("%+v\n", variable.O)
	// fmt.Printf("%+v\n", variable.C)
	// fmt.Printf("%+v\n", variable.H)
	// fmt.Printf("%+v\n", variable.L)
	// spew.Dump(variable.O)

	operator := []string{"-", "+", "*", "/", "(-)", "<=", "and", ">=", "<", ">", "or"}
	precedence := make(map[string]float32)
	setPrecedence(precedence)

	dojiFormula := node.DojiFormula
	postfixDojiExpression := generatePostfix(dojiFormula, operator, precedence)
	dojiObjectTree := generateObjectTree(postfixDojiExpression)
	generateJSONFromObject(dojiObjectTree, "doji.json")

	bullishEngulfingFormula := node.BullishEngulfing
	postfixBullishEngulfing := generatePostfix(bullishEngulfingFormula, operator, precedence)
	bullishEngulfingObjectTree := generateObjectTree(postfixBullishEngulfing)
	generateJSONFromObject(bullishEngulfingObjectTree, "bullishEngulfing.json")

	smaListFormula := node.SmaList
	postfixSmaList := generatePostfix(smaListFormula, operator, precedence)
	smaListObjectTree := generateObjectTree(postfixSmaList)
	generateJSONFromObject(smaListObjectTree, "smaList.json")

	emaListFormula := node.EmaList
	postfixEmaList := generatePostfix(emaListFormula, operator, precedence)
	emaListObjectTree := generateObjectTree(postfixEmaList)
	generateJSONFromObject(emaListObjectTree, "emaList.json")
}

func setPrecedence(precedence map[string]float32) {
	// https://www.geeksforgeeks.org/operator-precedence-and-associativity-in-c/
	precedence["*"] = 12.1
	precedence["/"] = 12.2
	precedence["%"] = 12.3

	precedence["+"] = 11.1
	precedence["-"] = 11.2

	precedence["<"] = 9.1
	precedence["<="] = 9.2
	precedence[">"] = 9.1
	precedence[">="] = 9.2

	precedence["=="] = 8.1
	precedence["!="] = 8.2

	precedence["and"] = 4 // Logical AND
	precedence["or"] = 3  // Logical OR

	precedence["(-)"] = 100
}

func generatePostfix(formula []any, operator []string, precedence map[string]float32) []any {
	var st stack
	ans := []any{}

	for i := 0; i < len(formula); i++ {
		currentString, isString := formula[i].(string)

		if currentString == "(" {
			st.push("(")
		} else if currentString == ")" {
			for !st.empty() && st.top() != "(" {
				ans = append(ans, st.top())
				st.pop()
			}
			st.pop()
		} else if isString && slices.Contains(operator, currentString) {
			for !st.empty() && precedence[currentString] <= precedence[st.top()] {
				ans = append(ans, st.top())
				st.pop()
			}
			st.push(currentString)
		} else {
			// spew.Dump(formula[i])
			ans = append(ans, formula[i])
		}
	}

	for !st.empty() {
		ans = append(ans, st.top())
		st.pop()
	}

	// got postfix
	// fmt.Println(len(ans))
	// for i := 0; i < len(ans); i++ {
	// spew.Dump(ans[i])
	// 	fmt.Printf("%+v \n", ans[i])
	// }
	return ans
}

func generateObjectTree(postfixExpression []any) any {
	st1 := []any{}
	for i := 0; i < len(postfixExpression); i++ {
		operatorString, isOperatorString := postfixExpression[i].(string)

		if isOperatorString && len(st1) >= 2 {
			node1 := st1[len(st1)-1].(node.Node)
			node2 := st1[len(st1)-2].(node.Node)

			st1 = st1[:len(st1)-2]
			var curNode node.Node
			curNode.Operator = operatorString
			curNode.Lhs = &node2
			curNode.Rhs = &node1

			st1 = append(st1, curNode)
		} else {
			st1 = append(st1, postfixExpression[i])
		}
	}

	if len(st1) != 1 {
		panic("generation failed")
	}

	return st1[0]
}

func generateJSONFromObject(objectTree any, fileName string) {
	generatedJSON, err := json.Marshal(objectTree)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(generatedJSON))
	fmt.Println("----")
	fmt.Println("----")
	fmt.Println("----")

	err = os.WriteFile(fileName, generatedJSON, 0644) // generates file in root project directory
	if err != nil {
		fmt.Print(err)
	}
}
