https://leetcode.com/problems/longest-common-prefix/description/
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

import "strings"
func longestCommonPrefix(strs []string) string {
    if len(strs)==0{
        return ""
    }
    if len(strs)==1{
        return strs[0]
    }
    // first sol, brute force
    ans := ""
    for i, v := range strs[0]{
        ans += string(v)
        fmt.Println(ans, i)
        var match bool
        for j:= 1; j<len(strs); j++{
            if strings.HasPrefix(strs[j],ans){
                match = true
            }else{
                match = false
                break
            }
        }
        if !match{
            return ans[:len(ans)-1]
        }
    }
    return ans
}