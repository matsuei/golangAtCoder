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
  
  x := 23
  // firstIndexみたいな感じ
  i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
  if i < len(data) && data[i] == x {
    // x is present at data[i]
  } else {
    // x is not present in data,
    // but i is the index where it would be inserted.
  }
}
