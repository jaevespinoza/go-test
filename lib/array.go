package lib

import "fmt"

func ArrayPlayground() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Len is %v\n", len(a))
	fmt.Printf("At space 5 is %v\n", a[4])

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
