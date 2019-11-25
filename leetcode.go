package main

import "log"

func main() {
	a := []int{2,7,11,15}
	v := 9
	log.Print(" 两数之和:",twoSum(a,v))
	log.Print(" 两数之和:",twoSum1(a,v))
}

func twoSum(nums []int, target int) []int {
	var res []int
	for i:=0;i<len(nums)-1;i++{
		def:=target-nums[i]
		for j:=i+1;j<len(nums);j++{
			if def==nums[j] {
				res=append(res,i)
				res=append(res,j)
				return res
			}
		}
	}
	return res
}

func twoSum1(nums []int, target int) []int {
	var res []int
	temp :=make(map[int]int)
	for i:=0;i<len(nums)-1;i++{
		temp[nums[i]]=i+1
		def:=target-nums[i]
		if temp[def]!=0 {
			res=append(res,temp[def]-1)
			res=append(res,i)
			return res
		}
	}
	return res
}