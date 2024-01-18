package lib

import "fmt"

type person struct {
	name string
	age  int
}

func StructPlayground(name string) {
	p := person{name: name}
	p.age = 42

	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 25}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp)
	sp.age = 45
	fmt.Println(sp)
	fmt.Println(s)
}
