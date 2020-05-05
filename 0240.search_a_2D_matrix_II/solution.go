package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix) - 1
	n := 0
	for m >= 0 && n < len(matrix[0]) {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			m--
		} else {
			n++
		}
	}
	return false
}

func main() {
	m1 := []int{1, 5, 9, 13, 17}
	m2 := []int{2, 6, 10, 14, 18}
	m3 := []int{3, 7, 11, 15, 19}
	m4 := []int{4, 8, 12, 16, 20}
	matrix := [][]int{m1, m2, m3, m4}
	print(searchMatrix(matrix, 15))
	print(searchMatrix(matrix, 21))
}
