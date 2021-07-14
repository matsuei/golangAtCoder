import (
  "bufio"
  "fmt"
)

const (
  initialBufSize = 10000
  maxBufSize     = 1000000
  mod            = 1e9 + 7
)

var (
  sc *bufio.Scanner = func() *bufio.Scanner {
    sc := bufio.NewScanner(os.Stdin)
    buf := make([]byte, initialBufSize)
    sc.Buffer(buf, maxBufSize)
    return sc
  }()
)

func nextStr() string {
  sc.Scan()
  return sc.Text()
}
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
  N := nextInt()
  L := make([]int, N)
  for i := 0; i < N; i++ {
    L[i] = nextInt()
  }
}
