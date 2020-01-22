// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "strings"
    "bytes"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    buf := new(bytes.Buffer)
    nbytes, err := buf.ReadFrom(resp.Body)
    body := buf.String()
    resp.Body.Close() // don't leak resources

	if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()

    filename := strings.Split(url, "//")[1] // remove http(s)://
    f, err := os.Create(filename)

	if err != nil {
        ch <- fmt.Sprintf("error creating file %s", filename)
        f.Close()
        return
    }

    intbytes, err := f.WriteString(body)
	if err != nil {
        ch <- fmt.Sprintf("while writing [%d] %s: %v", intbytes, filename, err)
        f.Close()
        return
    }
    f.Close()

    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
