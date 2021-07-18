var n int
fmt.Scan(&n)
var st string
fmt.Scan(&st)
slice := strings.Split(st, "")
x := make([]int, n)
for i := 0; i < n; i++ {
  in, e := strconv.Atoi(slice[i])
  if e != nil {
    panic(e)
  }
  x[i] = in
}
