array := make([]int, k)
m := make(map[int]bool)
uniq := [] int{}
for _, ele := range array {
  if !m[ele] {
    m[ele] = true
    uniq = append(uniq, ele)
  }
}

// 空のstructを使うパターン
m := make(map[int]struct{})

for _, ele := range arr {
  m[ele] = struct{}{} // m["a"] = struct{}{} が二度目は同じものとみなされて重複が消える。
}

uniq := [] int{}
for i := range m {
  uniq := append(uniq, i)
}
