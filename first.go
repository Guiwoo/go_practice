package main

import "fmt"

func shuffle(nums []int, n int) []int {
	rst := make([]int, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		rst[2*i] = nums[i]
		rst[2*i+1] = nums[i+n]
	}
	return rst
}

func findAnagrams(s string, p string) []int {
	arr := [26]int{}
	for _,v := range p {
		arr[v-'a']--
	}
	for i := 0; i < ; i++ {
		
	}
}

func zeroCheck(arr []int) bool {
	for _,v := range arr {
		if v < 0 return false
	}
	return true
}

func main() {
	ex := []int{2, 5, 1, 3, 4, 7}
	rst := shuffle(ex, 3)
	fmt.Println(rst)
}
