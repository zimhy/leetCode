package main

import (
	"container/list"
	"fmt"
)

/**
* basic calculator
 */

const (
	STATE_BEGIN    = 0
	STATE_NUMBER   = 1
	STATE_OPERATOR = 2
)

func compute(operators *Stack, numbers *Stack, back bool) {
	for {
		if numbers.Empty() || numbers.Len() < 2 || operators.Empty() {
			break
		}
		operator := operators.Pop()

		if '-' == operator {
			right := numbers.Pop()
			left := numbers.Pop()
			result := left.(int32) - right.(int32)
			numbers.Push(result)
		} else if '+' == operator {
			right := numbers.Pop()
			left := numbers.Pop()
			result := left.(int32) + right.(int32)
			numbers.Push(result)
		} else if '(' == operator {
			if !back {
				operators.Push(operator)
			}
			return
		}

	}

}

func calculate(s string) int {
	numberStack := NewStack()
	operatorStack := NewStack()
	runeArr := []rune(s)
	var number int32 = 0
	state := STATE_BEGIN
	firstRune := runeArr[0]
	if firstRune >= '0' && firstRune <= '9' {
		state = STATE_NUMBER
	} else {
		state = STATE_OPERATOR
	}
	for i := 0; i < len(runeArr); i++ {
		v := runeArr[i]
		if v == ' ' {
			continue
		}
		switch state {
		case STATE_NUMBER:
			if isNumber(v) {
				number = number*10 + (v - '0')
			} else {
				state = STATE_OPERATOR
				numberStack.Push(number)
				number = 0
				i--
			}
			break
		case STATE_OPERATOR:
			if isNumber(v) {
				state = STATE_NUMBER
				i--
				break
			}

			if v == '+' || v == '-' {
				compute(operatorStack, numberStack, false)
				operatorStack.Push(v)
				break
			}

			if '(' == v {
				operatorStack.Push(v)
			}
			if ')' == v {
				compute(operatorStack, numberStack, true)
			}
		}

	}
	if state == STATE_NUMBER {
		numberStack.Push(number)
	}

	compute(operatorStack, numberStack, false)
	result := numberStack.Pop().(int32)
	return int(result)
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
func main() {
	//data := "2-4-(8+2-6+(8+4-(1)+8-10))"
	data := "2+4-(8+10)"
	//data := "2-4-(8+2-6+(8+4-(1)+8-10))"
	fmt.Println("%s", calculate(data))
	//[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
	//var tickets = [][]string{
	//	[]string{"MUC","LHR"},
	//	[]string{"JFK","MUC"},
	//	[]string{"SFO","SJC"},
	//	[]string{"LHR","SFO"},
	//	//[]string{"MUC", "ABC"},
	//	//[]string{"LHR", "SFO"},
	//}

}
