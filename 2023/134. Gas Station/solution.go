package main

func main() {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}
	print(canCompleteCircuit(gas[:], cost[:]))
}
func canCompleteCircuit(gas []int, cost []int) int {
	start := -1
	tank := 0

	for i := 0; i < len(gas); i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			start = -1
			tank = 0
		} else if start < 0 {
			start = i
		}
	}

	for i := 0; i < start; i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			return -1
		}
	}
	return start
}
