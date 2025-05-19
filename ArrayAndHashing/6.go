https://leetcode.com/problems/group-anagrams/description/
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:

There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
Example 2:

Input: strs = [""]

Output: [[""]]

import "slices"
import "sort"

func isAnagram(s, r string) bool{
    t1:= strings.Split(s, "")
    t2:= strings.Split(r, "")
    sort.Strings(t1)
    sort.Strings(t2)
  //  fmt.Println(t1, t2)
    if slices.Equal(t1, t2){
        
        return true
    }
    return false
}
func groupAnagrams(strs []string) [][]string {
	// this is bruteforce approach
    ans := [][]string{}
    if len(strs)==0{
        return ans
    }
    if len(strs)==1{
        ans = append(ans,strs)
        return ans
    }
    // there could be some element already mapped to a group so we don't need to form another group for them.
    alreadyVisited := []int{}
    for i:=0; i<len(strs); i++{
        if slices.Contains(alreadyVisited, i){
            continue
        }
        ans = append(ans, []string{strs[i]})
     //   fmt.Println(i, strs[i], "i")
        for j:= i+1; j<len(strs); j++{
            if slices.Contains(alreadyVisited, j){
                continue
            }
            // check the possibility of anagram in other items
            // if anagram
           // fmt.Println(j, strs[j], "j")
            if isAnagram(strs[i], strs[j]){
                alreadyVisited = append(alreadyVisited, j)
                ans[len(ans)-1]= append(ans[len(ans)-1], strs[j])
            }
        }
    }
    return ans
}

