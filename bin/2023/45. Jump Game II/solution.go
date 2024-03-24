package main

func main() {
	print(jump([]int{2, 3, 0, 1, 4}))
}

func jump(nums []int) int {
	curJump, farthestJump, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthestJump {
			farthestJump = i + nums[i]
		}
 
		if i == curJump {
			jumps, curJump = jumps+1, farthestJump

			if curJump >= len(nums)-1 {
				return jumps
			}
		}
	}
 
	return 0
}