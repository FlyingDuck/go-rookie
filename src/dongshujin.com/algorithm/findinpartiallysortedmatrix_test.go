package algorithm

import "testing"
import "fmt"

func TestFindInpartiallySortedMatrix(t *testing.T) {

	array := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	found := FindInpartiallySortedMatrix(array, 7)
	fmt.Println("Found : ", found)
}
