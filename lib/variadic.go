package lib

import "fmt"

func VariadicPlayground(nums ...int) {
	total := 0

	for _, n := range nums {
		total += n
	}

	fmt.Println(total)
}
