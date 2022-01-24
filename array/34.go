package array

func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	left := findFirstPos(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLastPos(nums, target)
	return []int{left, right}
}
func findFirstPos(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
func findLastPos(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] == target {
			left = mid +1
		} else {
			right = mid - 1
		}
	}

	return right
}
