package main

import (
	"fmt"
	"strings"
)

//Problem 6 – Substitution Cipher
type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = s.name
	// your code here
	r := strings.NewReplacer(
		"a", "z",
		"b", "y",
		"c", "x",
		"d", "w",
		"e", "v",
		"f", "u",
		"g", "t",
		"h", "s",
		"i", "r",
		"j", "q",
		"k", "p",
		"l", "o",
		"m", "n",
		"n", "m",
		"o", "l",
		"p", "k",
		"q", "j",
		"r", "i",
		"s", "h",
		"t", "g",
		"u", "f",
		"v", "e",
		"w", "d",
		"x", "c",
		"y", "b",
		"z", "a")
	nameEncode = r.Replace(nameEncode)
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = s.nameEncode
	// your code here
	r := strings.NewReplacer(
		"z", "a",
		"y", "b",
		"x", "c",
		"w", "d",
		"v", "e",
		"u", "f",
		"t", "g",
		"s", "h",
		"r", "i",
		"q", "j",
		"p", "k",
		"o", "l",
		"n", "m",
		"m", "n",
		"l", "o",
		"k", "p",
		"j", "q",
		"i", "r",
		"h", "s",
		"g", "t",
		"f", "u",
		"e", "v",
		"d", "w",
		"c", "x",
		"b", "y",
		"a", "z")
	nameDecode = r.Replace(s.nameEncode)
	return nameDecode

}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student’s Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student’s Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
