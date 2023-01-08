package main

func main() {
	nums := [][]int{{1, 1}, {2, 2}, {3, 3}}
	print(maxPoints(nums[:]))
}
func maxPoints(points [][]int) int {
	// Map to store the slope of each line formed by a point and all other points
	slopeCount := make(map[float64]int)

	// Result variable to store the maximum number of points on a line
	var result int

	// Iterate over each point
	for i, p1 := range points {
		// Reset the slopeCount map for each point
		slopeCount = make(map[float64]int)

		// Set the maximum number of points on a line to 1 (the current point itself)
		maxPointsOnLine := 1

		// Set the number of overlapping points to 0
		overlap := 0

		// Iterate over the remaining points
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			// Check if the points are the same
			if p1[0] == p2[0] && p1[1] == p2[1] {
				overlap++
				continue
			}

			// Calculate the slope of the line formed by the two points
			slope := float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])

			// If the slope is infinite, set it to a large value
			if p2[0] == p1[0] {
				slope = 1e6
			}

			// Increment the count of points with this slope
			slopeCount[slope]++
		}

		// Find the maximum number of points on a line that includes the current point
		for _, count := range slopeCount {
			if count+1 > maxPointsOnLine {
				maxPointsOnLine = count + 1
			}
		}

		// Add the overlap to the maximum number of points on a line
		maxPointsOnLine += overlap

		// Update the result if necessary
		if maxPointsOnLine > result {
			result = maxPointsOnLine
		}
	}

	return result
}
