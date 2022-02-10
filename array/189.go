package array

func rotate(nums []int, k int) {
	step := k % len(nums)
	rotateSlice(nums)
	rotateSlice(nums[:step])
	rotateSlice(nums[step:])
}

func rotateSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]

	}
}
