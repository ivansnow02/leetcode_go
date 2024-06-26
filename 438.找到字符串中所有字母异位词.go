package main
/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (53.57%)
 * Likes:    1429
 * Dislikes: 0
 * Total Accepted:    422.8K
 * Total Submissions: 790.7K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 * 
 * 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: s = "cbaebabacd", p = "abc"
 * 输出: [0,6]
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "abab", p = "ab"
 * 输出: [0,1,2]
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= s.length, p.length <= 3 * 10^4
 * s 和 p 仅包含小写字母
 * 
 * 
 */

// @lc code=start
// func findAnagrams(s string, p string) []int {
// 	var res []int
// 	pCode := encodeS(p)
// 	for i := 0; i < len(s)-len(p)+1; i++ {
// 		sCode := encodeS(s[i:i+len(p)])

// 		if sCode == pCode {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

// func encodeS(s string) int {
// 	primes := []int{2, 3, 5, 7, 11, 13, 17, 19,
// 		23, 29, 31, 37, 41, 43, 47, 53,
// 		59, 61, 67, 71, 73, 79, 83, 89,
// 		97, 101}
// 	ans := 1
// 	for _, b := range s {
// 		ans *= primes[b-'a']
// 		ans %= 1e9 + 7
// 	}
// 	return ans
// }

func findAnagrams(s string, t string) []int {
    need := make(map[byte]int)
    window := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        need[t[i]]++
    }

    left, right := 0, 0
    valid := 0
    var res []int
    for right < len(s) {
        c := s[right]
        right++
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        for right - left >= len(t) {
            if valid == len(need) {
                res = append(res, left)
            }
            d := s[left]
            left++
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    return res
}
// @lc code=end

