package assignments

import "fmt"

func main() {
	x := 1

	v := 1
	v++
	v--

	x, v = v, x

	medals := []string{"gold", "silver", "bronze"}
	// same as
	medals[0] = "gold"
	medals[1] = "silver"
	medals[2] = "bronze"

	
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}

