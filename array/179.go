package array

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	ans := []byte{}
	for _, num := range nums {
		ans = append(ans, strconv.Itoa(num)...)
	}
	return string(ans)
}
