package main

import (
	"fmt"
	"math"
)

const st = "the root square is: "

func switchGo(i int) {
	switch i {
	case 1:
		fmt.Println("i is one")
	case 2, 3:
		fmt.Println("i is not one")
	default:
		fmt.Println("i is far away")
	}

	whatAmI := func(i interface{}) {
		switch tys := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		default:
			fmt.Printf("Don't know type %T\n", tys)
		}
	}

	whatAmI(true)
	whatAmI(2)
}

func arrayPlayground() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Len is %v\n", len(a))
	fmt.Printf("At space 5 is %v\n", a[4])

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func main() {
	x := 8.0
	fmt.Println(math.Sqrt(x))
	fmt.Println(st)

	if x == 8 || x < 8 {
		fmt.Println("New prize")
	} else {
		fmt.Println("No prize")
	}

	switchGo(2)
	arrayPlayground()
}
