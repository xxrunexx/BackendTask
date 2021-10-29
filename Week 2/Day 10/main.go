package main

import "fmt"

type alphaMap map[rune]int

func Chara(s string) alphaMap {
	m := alphaMap{}
	for _, v := range s {
		m[v]++
	}
	return m
}

func GetFreqList(val string) {
	character := make(chan alphaMap)
	go func(text string) {
		character <- Chara(text)
	}(val)
	for char, count := range <-character {
		if string(char) != "," && string(char) != " " {
			fmt.Println(string(char), ":", count)
		}
	}
}

func main() {
	GetFreqList("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
}
