package dynamic

import (
	"fmt"
	"testing"
)

//public int countPaths(boolean[][] grid, int row, int col) {
//    int[][] opt = new int[row][col];
//    for (int r = row - 1; r >= 0; r--) {
//        for (int c = col - 1; c >= 0; c--) {
//            if ((r == row - 1) || (c == col - 1)) {
//                opt[r][c] = 1;
//            } else if (grid[r][c]) {
//                opt[r][c] = 0;
//            } else if (!grid[r][c]) {
//                opt[r][c] = opt[r + 1][c] + opt[r][c + 1];
//            }
//        }
//    }
//    return opt[0][0];
//}
func TestMigong(t *testing.T) {
	var mat = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}

	anser(mat)
}

func anser(mat [][]int) {

	var opt = [][]int{}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			if mat[i][j] == 1 {
				opt[i][j] = 0
			} else {
				opt[i][j] = opt[i-1][j] + opt[i][j-1]
			}
		}
	}

	fmt.Println(opt)
}
