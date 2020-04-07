package new

import "fmt"

func newInt1() *int {
    return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	p := newInt1()
	q := newInt2()

	fmt.Println(p == q) // "false"
	
}