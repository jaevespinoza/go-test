package lib

import (
	"fmt"
	"slices"
)

func SlicePlayground() {
	s := make([]string, 3)
	fmt.Println("emp: ", s, "len:", len(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "f", "h")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	sli := c[1:2]
	fmt.Println(sli)

	if slices.Equal(sli, c) {
		fmt.Println("equals!")
	} else {
		fmt.Println("not equals!")
	}

	multiD := make([][]string, 3)
	multiD[1] = []string{"1", "2", "3"}

	fmt.Println(multiD)
}
