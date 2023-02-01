package main

import "fmt"

func main() {
	sum := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(sum)
}

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
func twoSum(nums []int, target int) []int {
	//先定义一个容器池
	indexAndValue := make(map[int]int)
	//循环题目上的数组
	for index, value := range nums {
		//判断容器是否包含
		if indexByNeedValue, ok := indexAndValue[target-value]; ok {
			//进来就是包含了,我们直接返回内容
			return []int{indexByNeedValue, index}
		}
		//到这里就是还没有,我们存放需要的
		indexAndValue[value] = index
	}
	return nil
}
