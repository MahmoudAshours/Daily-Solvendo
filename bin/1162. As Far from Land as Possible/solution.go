package main

func main(){
print(maxDistance([][]int{{1,0,1},{0,0,0},{1,0,1}}))
}
func maxDistance(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	MAX_DISTANCE := rows + cols + 1

	// First pass: check for left and top neighbors
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				// Distance of land cells will be 0.
				grid[i][j] = 0
			} else {
				grid[i][j] = MAX_DISTANCE
				// Check left and top cell distances if 
they exist and update the current cell distance.
				if i > 0 {
					grid[i][j] = min(grid[i][j], 
grid[i-1][j]+1)
				}
				if j > 0 {
					grid[i][j] = min(grid[i][j], 
grid[i][j-1]+1)
				}
			}
		}
	}

	// Second pass: check for the bottom and right neighbors.
	ans := -2147483648
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			// Check the right and bottom cell distances if 
they exist and update the current cell distance.
			if i < rows-1 {
				grid[i][j] = min(grid[i][j], 
grid[i+1][j]+1)
			}
			if j < cols-1 {
				grid[i][j] = min(grid[i][j], 
grid[i][j+1]+1)
			}

			ans = max(ans, grid[i][j])
		}
	}

	// If ans is 0, it means there is no water cell,
	// If ans is MAX_DISTANCE, it implies no land cell.
	if ans == 0 || ans == MAX_DISTANCE {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
