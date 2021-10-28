package main

import "fmt"

func SimpleEquations(a, b, c int) {
	var length int
	if c < a || c < b {
		fmt.Println("No Solution")
		return
	}
	if a > b {
		length = a
	} else {
		length = b
	}
	if length <= 5 {
		fmt.Println("No Solution")
		return
	}
	for i := 1; i <= (length/2)-2; i++ {
		if i+(i+1)+(i+2) == a && i*(i+1)*(i+2) == b && (i*i)+((i+1)*(i+1))+((i+2)*(i+2)) == c {
			fmt.Println(i, i+1, i+2)
			return
		}
		for j := i + 1; j <= (length/2)-1; j++ {
			if i+j+(j+1) == a && i*j*(j+1) == b && (i*i)+(j*j)+((j+1)*(j+1)) == c {
				fmt.Println(i, j, j+1)
				return
			}
			for k := j + 1; k <= (length / 2); k++ {
				if i+(j)+(k) == a && i*(j)*(k) == b && (i*i)+((j)*(j))+((k)*(k)) == c {
					fmt.Println(i, j, k)
					return
				}
			}
		}
	}
	fmt.Println("No Solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
