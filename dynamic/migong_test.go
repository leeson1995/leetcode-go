package dynamic

import (
	"fmt"
	"testing"
)

func TestMigong(t *testing.T) {
	var mat = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}

	fmt.Println(uniquePathsWithObstacles(mat))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
				obstacleGrid[i][j] = 1
			}
			if i == 0 && j != 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				}
			}
			if i != 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				}
			}
			if i != 0 && j != 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	fmt.Println(obstacleGrid)
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
