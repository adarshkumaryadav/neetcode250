https://leetcode.com/problems/two-sum/description/
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(nums []int, target int) []int {
	// 1st sol
    // brute force 2 iteration. i=o <len-1 & j=i+1 j<len, add and return i and j
    // o (n*n)

    // 2nd approach
    // using map and a for loop
    tempMap := make(map[int]int)
    // store the value as key and index as value
    for i:= 0; i< len(nums); i++{
        remaining := target - nums[i]
        // see if the map has the remaining no. as the key
        if v, ok := tempMap[remaining]; ok{
            return []int{i, v}
        }
        // if not present store this value
        tempMap[nums[i]]=i
    }
    return []int{}
}