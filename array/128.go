package array

// 128.最长连续序列【并查集】

func find(h []int, x int) int {
	if h[x] == x {
		return x
	}
	return find(h, h[x])
}
func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums))
	h := make([]int, len(nums))
	counts := make([]int, len(nums))
	max := 0
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			continue
		}
		m[nums[i]] = i
		h[i] = i
		counts[i] = 1
		left, ok := m[nums[i]-1]
		if ok {
			head := find(h, left)
			h[head] = i
			counts[i] = counts[head] + counts[i]
		}

		right, ok := m[nums[i]+1]
		if ok {
			head := find(h, right)
			h[head] = i
			counts[i] = counts[head] + counts[i]
		}

		if max < counts[i]{
			max = counts[i]
		}
	}
	return max
}
