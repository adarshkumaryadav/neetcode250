// https://leetcode.com/problems/contains-duplicate/description/
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct. 

Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.
// import "sort"
func containsDuplicate(nums []int) bool {
	/* first sol
	 // create a map ds
	 // map is a referenced type so, need to initialize it with make
	 // o(n)
	 tempMap := make(map[int]bool)
	 for _, v := range nums{
		 if _, ok := tempMap[v]; ok{
			 return true
		 }
		 tempMap[v]= true
	 }
	 return false
	 */
	 // second solution: brute force
	 // time complexity = o(n*n)
	 // space complexity = o(1)
	 /*
	 for i:=0; i<len(nums); i++{
		 for j:= i+1; j<len(nums); j++{
			 if nums[i]==nums[j]{
				 return true
			 }
		 }
	 }
	 return false
	 */
	 // 3rd solution
	 // sorting package
	 // n + n log n = log n
	 /*
	 if len(nums) <= 1{
		 return false
	 }
	 // O(n log n)
	 // space log n
	 sort.Ints(nums)
	 for i:= 0; i<len(nums)-1; i++{
		 if nums[i]==nums[i+1]{
			 return true
		 }
	 }
	 return false
	 */
	 // multiset
	 // map[num]occurance // occurance is int
	 s:= newMultiSet()    
	 s.addNums(nums)
	 for _, v := range nums{
		 if s.hasDuplicate(v){
			 return true
		 }
	 }
	 return false
 
 }
 // create a multiset
 type multiSet struct{
	 items map[int]int
 }
 
 func newMultiSet() *multiSet{
	 return &multiSet{
		 items: make(map[int]int),
	 }
 }
 func (s *multiSet)addNums(nums []int){
	 for _, num := range nums{
		 s.addNewElement(num)
	 }
 }
 func (s *multiSet)addNewElement(num int){
	 s.items[num]++
 }
 func (s *multiSet)hasDuplicate(n int) bool{
	 if v, ok := s.items[n]; ok && v > 1{
		 return true
	 }
	 return false
 }