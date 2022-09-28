package leetcode

import "sort"

//https://leetcode.com/problems/two-sum/
type Num struct {
	value int
	index int
}

func twoSum(nums []int, target int) []int {

	list := make([]Num, len(nums))

	for i, v := range nums {
		list[i] = Num{v, i}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].value < list[j].value
	})

	for _, v := range list {
		rs := sort.Search(len(list), func(i int) bool {
			return list[i].value >= target-v.value
		})

		if rs != len(list) && v.index != list[rs].index && list[rs].value == target-v.value {
			return []int{v.index, list[rs].index}
		}
	}

	return []int{}
}
