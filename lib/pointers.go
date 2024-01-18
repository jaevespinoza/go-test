package lib

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func PointerPlayground() {
	i := 22

	fmt.Println("initial: ", i)
	// Distinct copy of the initial value
	zeroval(i)
	fmt.Println("zeroval: ", i)

	// replaces the memory value instead of the copy
	// zeroval doesn’t change the i in main,
	// but zeroptr does because it has a reference
	// to the memory address for that variable.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
}
