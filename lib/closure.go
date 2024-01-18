package lib

func ClosurePlayground() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
