https://leetcode.com/problems/majority-element/
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
Example 1:

Input: nums = [3,2,3]
Output: 3

func majorityElement(nums []int) int {
    tempMap := make(map[int]int)
    maxCond:= len(nums)/2
    for _, v:= range nums{
        tempMap[v]++
        if tempMap[v]>maxCond{
            return v
        }
    }
    /*for k, v := range tempMap{
        if v>len(nums)/2{
            return k
        }
    }
    */
    return 0
}