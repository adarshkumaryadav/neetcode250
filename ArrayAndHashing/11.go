https://leetcode.com/problems/sort-an-array/
o(n logn)
Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.


func sortArray(nums []int) []int {
    return mergeSort(nums)
}
func mergeSort(nums []int)[]int{
    // check is the array has just an element or 0 element
    if len(nums)<2{
        return nums
    }
    // divide
    mid := len(nums)/2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])

    // conquer : adding the sol 
    return merge(left, right)
}

func merge(left []int, right []int)[]int{  
    x, y := len(left), len(right)
    temp := make([]int,0)
    i, j := 0, 0
    for i < x && j<y {
        if left[i]<right[j]{
            temp = append(temp, left[i])
            i++
        }else{
            temp = append(temp, right[j])
            j++
        }
    }
    if i < x{
        temp = append(temp, left[i:]...)
    }
    if j < y{
        temp = append(temp, right[j:]...)
    }
    return temp
}