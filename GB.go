package main



func SubSort(left,right []int)[]int{
	result := make([]int,0)
	i,j := 0,0
	l,r := len(left),len(right)
	for i < l && j > r{
		if left[i] > right[j]{
			result =  append(result,right[j])
			j++
		}else {
			result = append(result,left[i])
			i++
		}
	}
	result = append(result,right[j:]...)
	result = append(result,left[i:]...)
	return result
}

func Sort(nums []int)[]int{
	if len(nums) < 2{
		return nums
	}
	i := len(nums)/2
	left := Sort(nums[0:i])
	right := Sort(nums[i:])
	result := SubSort(left,right)
	return result
}

