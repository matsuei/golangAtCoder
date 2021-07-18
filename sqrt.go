package main

import (
  "math"
)

func main() {
  sc.Split(bufio.ScanWords)
  n := nextInt()
  max := int(math.Sqrt(float64(n)))
  var slice []int
  for i := 1; i < max; i++ {
    if n % i == 0 {
      slice = append(slice, i, n / i)
    }
  }
  sort.Sort(sort.IntSlice(slice))
  for i := 0; i < len(slice); i++ {
    fmt.Printf("%d\n", slice[i])
  }
}
