package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	// Convert String to Array using Strings lib
	convArr := strings.Split(angka, "")

	hasilAppend := []int{}
	hasilFilter := [10]int{}

	for _, i := range convArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Print("nil")
		}
		if hasilFilter[j] == 10 {
			continue
		} else if hasilFilter[j] == 0 {
			hasilFilter[j] = j
		} else if hasilFilter[j] == j {
			hasilFilter[j] = 10
		}

		for _, e := range hasilFilter {
			if e != 0 && e != 10 {
				hasilAppend = append(hasilAppend, e)
			}
		}
	}
	return hasilAppend
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
