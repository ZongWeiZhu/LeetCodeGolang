package array

// 快排
func sortArray(nums []int) []int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	return nums
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	counter := start
	pivot := end
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[counter] = nums[counter], nums[i]
			counter++
		}
	}
	nums[pivot], nums[counter] = nums[counter], nums[pivot]
	quickSort(nums, start, counter-1)
	quickSort(nums, counter+1, end)
}

// 归并排序
func sortArray(s []int) []int{
	len := len(s)
	if len == 1 {
		return s //最后切割只剩下一个元素
	}
	m := len/2
	leftS := sortArray(s[:m])
	rightS := sortArray(s[m:])
	return merge(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func merge(l []int, r []int) []int{
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0, len(l)+len(r))

	lIndex,rIndex := 0,0 //两个切片的下标，插入一个数据，下标加一
	for lIndex<lLen && rIndex<rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		}else{
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}