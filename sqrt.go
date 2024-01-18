package main

import (
	"example/playground/lib"
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

func main() {
	x := 8.0
	fmt.Println(math.Sqrt(x))
	fmt.Println(st)

	if x == 8 || x < 8 {
		fmt.Println("New prize")
	} else {
		fmt.Println("No prize")
	}

	// switchGo(2)
	// lib.ArrayPlayground()
	// lib.SlicePlayground()
	// lib.MapsPlayground()
	// lib.RangePlayground()
	// fmt.Println(lib.Multifunction(4, 5))
	// lib.VariadicPlayground()
	// fmt.Println(lib.RecursionPlayground())
	// lib.PointerPlayground()
	// lib.StringPlayground()
	lib.StructPlayground("name")
}
