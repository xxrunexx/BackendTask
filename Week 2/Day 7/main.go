package main

import (
	"fmt"
	"regexp"
	"sort"
)

func factorial(n int) int { //Fungsi Rekursif
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

type stack []int

func (s *stack) Pop() int {
	num := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return num
}

func (s *stack) Push(num int) {
	*s = append((*s), num)
}

type queue []int

func (q *queue) Pop() int {
	num := (*q)[0]
	*q = (*q)[1:]
	return num
}

func (q *queue) Push(num int) {
	*q = append((*q), num)
}

func main() {

	var a []string

	a = append(a, "string1")
	a = append(a, "string2")
	a = append(a, "string10")
	a = append(a, "string3")
	a = append(a, "string6")
	a = append(a, "string4")

	sort.Strings(a)
	idx := sort.SearchStrings(a, "string10")

	fmt.Printf("array : %+v \n index : %d\n", a, idx)

	fmt.Println(factorial(5))

	var stack1 stack

	stack1.Push(10)
	stack1.Push(15)
	stack1.Push(1)

	fmt.Println(stack1)

	fmt.Println(stack1.Pop())
	fmt.Println(stack1.Pop())

	fmt.Println(stack1)

	stack1.Push(4)

	fmt.Println(stack1)

	fmt.Println(stack1.Pop())

	fmt.Println(stack1)

	var queue1 queue

	queue1.Push(10)
	queue1.Push(1)
	queue1.Push(5)

	fmt.Println(queue1)

	fmt.Println(queue1.Pop())
	fmt.Println(queue1.Pop())

	fmt.Println(queue1)

	stringEmail := "kamil@alterra"
	regex, _ := regexp.Compile(`\S+@\S+\.\S+`)

	fmt.Println(regex.Match([]byte(stringEmail)))

}
