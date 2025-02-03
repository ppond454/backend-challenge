package test1

import (
	"encoding/json"
	"fmt"
	"os"
)

type IGraph[T Number] interface {
	findRoute()
	max(a T, b T) T
	findRouteMax2()
}

type Number interface {
	int8 | int | int64 | float32 | float64
}

type Graph[T Number] struct {
	data [][]T
}

func (g *Graph[T]) findMax() T {
	var clone [][]T
	data, _ := json.Marshal(g.data)
	json.Unmarshal(data, &clone)

	cursor := len(clone) - 1
	for 0 < cursor {
		for idx := range cursor {
			max := g.max(clone[cursor][idx], clone[cursor][idx+1])
			next := &clone[cursor-1][idx]
			*next += max
		}
		cursor -= 1
	}

	return clone[0][0]
}

func (g *Graph[T]) max(a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// var dataTemp = [][]int{
// 	{59},
// 	{73, 41},
// 	{52, 40, 53},
// 	{26, 53, 6, 34},
// }

func Run() {
	jsonData, err := os.ReadFile("./files/hard.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data [][]int
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	g := &Graph[int]{data: data}
	max := g.findMax()
	fmt.Println("max sum:", max)

}
