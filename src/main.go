package main

// 1. Install bundle go get -u golang.org/x/tools/...
//go:generate bundle -o ./../out/bundle.go -prefix "" .

func SomeLinq() int {
	a := []int{2, 4, 5, 6, 7, 11, 1}
	biggest := From(a).OrderByDescending(
		func(i interface{}) interface{} {
			return i.(int)
		}).First()
	return biggest.(int)
}

/**
 * Copy main function from puzzle template
 **/
func main() {
	debugln("main!")
}
