package main

import (
	"fmt"
	"os"
)

// 读取文件中的迷宫
func readMaze() [][]int {
	filename := "learnGo/chapter_12/maze.in"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) // 使用Fscanf读取文件中第一行的值

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j]) // 使用Fscanf读取文件中的迷宫
		}
	}
	return maze
}

type point struct { // 抽象点
	i, j int
}

var dirs = [4]point{ // 增加对于一个点广度优先的顺序
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point { // 添加 点 + 点 的方法
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) { // 添加判断点是否在迷宫内
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start} // 定义探测点的储存队列
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end { // 判断结束
			break
		}

		for _, dir := range dirs { // 遍历当前点周围的点
			next := cur.add(dir)
			// 该元素在 maze 中，没有墙 maze(next) == 0,
			// 而且没有被探测过 steps(next) == 0 ,
			// 而且不等于start steps(next) != start,
			val, ok := next.at(maze) // 读取探测点在迷宫中的位置
			if !ok || val == 1 {     // 当前点不在迷宫中 或者 当前点是墙
				continue
			}

			val, ok = next.at(steps) // 探测点不在已探测的二维数组 或者  当前点已经被探测过
			if !ok || val != 0 {
				continue
			}

			if next == start { // 探测点不能是起点
				continue
			}
			curStepVal, _ := cur.at(steps)         // 获取当前点的步数
			steps[next.i][next.j] = curStepVal + 1 // 修改探测点的步数值
			Q = append(Q, next)                    // 添加下次探测点的到队列中
		}
	}
	return steps
}

func main() {

	maze := readMaze() // 读取文件中迷宫信息
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1}) // 获取探测结果
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}
