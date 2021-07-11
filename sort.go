package main

import (
  "fmt"
  "sort"
)

func main() {
  var n int
  fmt.Scan(&n)
  x := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&x[i])
  }
  sort.Sort(sort.IntSlice(x))
  sort.Sort(sort.Reverse(sort.IntSlice(x)))
}
