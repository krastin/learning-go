package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	names := make(map[string]string)
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, nil)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, names)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t%d\t%s\n", names[line], n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int, names map[string]string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] = f.Name()
    }
    // NOTE: ignoring potential errors from input.Err()
}