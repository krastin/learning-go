package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "exercise2: %v\n", err)
            os.Exit(1)
		}
		
		// convert f <-> c
        fmt.Printf("%.2f°F = %.2f°C, %.2f°C = %.2f°F\n", t, ((t-32)*5/9), t, ((t*9/5)+32))
		// convert ft <-> m
        fmt.Printf("%.2fft = %.2fm, %.2fm = %.2fft\n", t, t*0.3048, t, t/0.3048)
		// convert lbs <-> kg
		fmt.Printf("%.2flbs = %.2fkg, %.2fkg = %.2flbs\n", t, t/2.205, t, t*2.205)
    }
}