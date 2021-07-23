package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextString())
	return n
}

type Node struct {
	id    int
	link  []*Node
	check bool
}

func dfs(node *Node) int {
	res := 1
	node.check = true

	for _, next_node := range node.link {
		if !next_node.check {
			res += dfs(next_node)
		}
	}
	return res
}

func reset(nodes []*Node) {
	for _, node := range nodes {
		node.check = false
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, M := nextInt(), nextInt()

	nodes := make([]*Node, N+1)
	for i := 0; i < N+1; i++ {
		nodes[i] = &Node{}
		nodes[i].id = i
		nodes[i].check = false
	}

	for i := 0; i < M; i++ {
		a, b := nextInt(), nextInt()
    // 有向グラフの場合
		nodes[a].link = append(nodes[a].link, nodes[b])
    // 無向グラフの場合
    nodes[a].link = append(nodes[a].link, nodes[b])
    nodes[b].link = append(nodes[b].link, nodes[a])
	}

	ans := 0
	for i := 1; i < N+1; i++ {
		ans += dfs(nodes[i])
		reset(nodes)
	}
	fmt.Println(ans)
}

