package main


func flipAndInvertImage(A [][]int) [][]int {
	for _, i := range A {
		l := len(i)
		for j := 0; j < (l+1)/2; j++ {
			i[j], i[l-j-1] = i[l-j-1]^1, i[j]^1
		}
	}
	return A
}
