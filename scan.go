package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
    sc.Split(bufio.ScanWords)
    n := nextInt()
    x := make([]int, n)
    for i := 0; i < n; i++ {
        x[i] = nextInt()
    }
    fmt.Printf("%d", x)
}
