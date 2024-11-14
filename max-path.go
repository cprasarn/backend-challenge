package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Data [][]int `json:"data"`
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
	Sum   []int
}

func tree(in [][]int) *Node {
	root := Node{Value: in[0][0]}
	n := make([][]*Node, 1)
	n[0] = append(n[0], []*Node{&root}...)
	l := len(in)
	for i := 1; i < l; i++ {
		m := make([]*Node, 0)
		k := 0
		for j := 0; j < len(in[i]) && j+1 < len(in[i]) && k < len(n[i-1]); j++ {
			p := n[i-1][k]
			if j == 0 {
				left := Node{Value: in[i][j]}
				p.Left = &left
				m = append(m, p.Left)
			} else {
				p.Left = m[j]
				p.Left.Sum = append(p.Left.Sum)
			}

			right := Node{Value: in[i][j+1]}
			p.Right = &right
			m = append(m, p.Right)
			k++
		}
		n = append(n, [][]*Node{m}...)
	}
	return &root
}

func max(root *Node, sum int) int {
	maxResult := sum + root.Value
	leftResult := 0
	if root.Left != nil {
		leftResult = max(root.Left, maxResult)
	}

	rightResult := 0
	if root.Right != nil {
		rightResult = max(root.Right, maxResult)
	}

	if leftResult > rightResult {
		maxResult = leftResult
	} else if rightResult > 0 {
		maxResult = rightResult
	}

	return maxResult
}

func main() {
	f, err := os.Open("./files/hard.json")
	if err != nil {
		log.Fatalln(err)
	}

	buffer := []byte{}
	n := 1
	for n > 0 {
		bufferString := make([]byte, 1000)
		n, err := f.Read(bufferString)
		if err != nil {
			break
		}

		if n > 0 {
			buffer = append(buffer, bufferString[:n]...)
		}
	}

	s := string(buffer)
	data := Data{}
	if n > 0 {
		err = json.Unmarshal([]byte(s), &data)
		if err != nil {
			fmt.Println("Cannot unmarshal:", s)
			log.Fatalln(err)
		}
	}

	root := tree(data.Data)
	maxValue := max(root, 0)
	fmt.Println("Output:", maxValue)
}
