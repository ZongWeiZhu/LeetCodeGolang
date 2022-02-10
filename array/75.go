package array

/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

 

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]
 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortColors(nums []int) {
	p0, p1, p2 := 0, 0, len(nums)-1

	if p2 <= 0 {
		return
	}

	for p1 <= p2 {
		if nums[p1] == 2 {
			nums[p1] = nums[p2]
			nums[p2] = 2
			p2--
		} else if nums[p1] == 0 {
			nums[p1] = nums[p0]
			nums[p0] = 0
			p0++
			p1++
		} else {
			p1++
		}
	}
}
