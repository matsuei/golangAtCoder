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
    x := make([]int, n)
    x := 0
    for i := 0; i < n; i++ {
        x[i] = nextInt()
    }
    fmt.Println(x)
}
