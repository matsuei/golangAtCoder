package main

import (
  "fmt"
  "sort"
)
type Object struct {
  Index string
  Value  int
}

func main() {
  var n int
  fmt.Scan(&n)
  x := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&x[i])
  }
  sort.Sort(sort.IntSlice(x))
  sort.Sort(sort.Reverse(sort.IntSlice(x)))
  
  sort.Ints(index)
  
  sort.Slice(people, func(i, j int) bool { 
    return Object[i].Index < Object[j].Index
  })
}
