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

func main() {
  sc.Split(bufio.ScanWords)
  N := nextInt()
  L := make([]int, N)
  for i := 0; i < N; i++ {
    L[i] = nextInt()
  }
}
