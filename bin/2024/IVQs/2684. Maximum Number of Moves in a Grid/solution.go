package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxMoves([][]int{
		{1000000, 92910, 1021, 1022, 1023, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033, 1034, 1035, 1036, 1037, 1038, 1039, 1040, 1041, 1042, 1043, 1044, 1045, 1046, 1047, 1048, 1049, 1050, 1051, 1052, 1053, 1054, 1055, 1056, 1057, 1058, 1059, 1060, 1061, 1062, 1063, 1064, 1065, 1066, 1067, 1068},
		{1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079, 1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103, 1104, 1105, 1106, 1107, 1108, 1109, 1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118},
	}))
}

func maxMoves(grid [][]int) int {
	maxCount := 0
	minim := math.MaxInt
	minimIndex := 0

	// Find the minimum value in the first column
	for i := 0; i < len(grid); i++ {
		if grid[i][0] < minim {
			minim = grid[i][0]
			minimIndex = i
		}
	}
	recur(grid, minimIndex, 0, 0, &maxCount)
	return maxCount
}

func recur(grid [][]int, i, j, currentCount int, maxCount *int) {
	// Base case: check if indices are out of bounds
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	// Update maxCount if current path is longer
	if currentCount > *maxCount {
		*maxCount = currentCount
	}

	// Try moving in all three possible directions
	if i > 0 && j < len(grid[0])-1 && grid[i-1][j+1] > grid[i][j] {
		recur(grid, i-1, j+1, currentCount+1, maxCount)
	}
	if j < len(grid[0])-1 && grid[i][j+1] > grid[i][j] {
		recur(grid, i, j+1, currentCount+1, maxCount)
	}
	if i < len(grid)-1 && j < len(grid[0])-1 && grid[i+1][j+1] > grid[i][j] {
		recur(grid, i+1, j+1, currentCount+1, maxCount)
	}
}
