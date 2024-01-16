package lib

import "fmt"

func MapsPlayground() {
	m := make(map[string]int)
	m["k1"] = 2
	m["k2"] = 1

	fmt.Println(m)
	delete(m, "k1")

	fmt.Println(m)
	clear(m)

	_, inMap := m["k2"]
	fmt.Println(inMap)
}
