https://leetcode.com/problems/valid-anagram/description/
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:

Input: s = "anagram", t = "nagaram"

Output: true


import "sort"

//import "slices"
import "strings"

func isAnagram(s string, t string) bool {
	// first solution
	// if lengths are not equal, for sure they are not anagram
	if len(s) != len(t) {
		return false
	}
	/* // store the char in the map and increase the value
	   tempMap := make(map[rune]int)
	   for _, v := range s {
	       tempMap[v]++
	   }
	   // decrease the value
	   // if element is not found it will be - or if more than the count in s
	   for _, v := range t {
	       tempMap[v]--
	       if tempMap[v] < 0{
	           return false
	       }
	   }
	   return true
	*/

	// second approach using sorting

	temp1 := []rune(s)
	temp2 := []rune(t)
	sort.Slice(temp1, func(i, j int) bool { return temp1[i] < temp1[j] })
	sort.Slice(temp2, func(i, j int) bool { return temp2[i] < temp2[j] })
	// return slices.Equal(temp1, temp2)
	if string(temp1) != string(temp2) {
		return false
	}
	return true

	// third approach, using sorting
	/*  t1:= strings.Split(s, "")
	    t2:= strings.Split(t, "")
	    // above 2 are the slices of string
	    sort.Strings(t1)
	    sort.Strings(t2)
	    for i, v := range t1{
	        if t2[i]!=v{
	            return false
	        }
	    }
	    return true
	*/
}