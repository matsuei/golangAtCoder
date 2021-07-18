array := make([]int, k)
m := make(map[int]bool)
uniq := [] int{}
for _, ele := range array {
  if !m[ele] {
    m[ele] = true
    uniq = append(uniq, ele)
  }
}
