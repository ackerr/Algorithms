package leetcode0994

import "fmt"

type position struct {
	x, y, level int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	R, C := len(grid), len(grid[0])
	var queue []position
	for i, rows := range grid {
		for j, v := range rows {
			if v == 2 {
				queue = append(queue, position{x: i, y: j, level: 0})
			}
		}
	}

	ans := 0
	for len(queue) > 0 {
		pos := queue[0]
		i, j := pos.x, pos.y
		ans = pos.level
		s := []position{position{i - 1, j, 0}, position{i + 1, j, 0}, position{i, j - 1, 0}, position{i, j + 1, 0}}
		for _, v := range s {
			row, col := v.x, v.y
			if 0 <= row && row < R && 0 <= col && col < C && grid[row][col] == 1 {
				grid[row][col] = 2
				queue = append(queue, position{x: row, y: col, level: ans + 1})
			}
		}
		queue = queue[1:]
	}

	for _, rows := range grid {
		for _, v := range rows {
			if v == 1 {
				return -1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}
