package main
/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode.cn/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (57.56%)
 * Likes:    1038
 * Dislikes: 0
 * Total Accepted:    516.2K
 * Total Submissions: 896.7K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给你两个整数数组 nums1 和 nums2
 * ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2,2]
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[4,9]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 1000
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 * 
 * 
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	count1 := make([]int, 1001)
    count2 := make([]int, 1001)
    ans := make([]int, 0)
    for _, v := range nums1 {
        count1[v]++
    }
    for _, v := range nums2 {
        count2[v]++
    }
    for i := 0; i <= 1000; i++ {
        for count1[i] > 0 && count2[i] > 0 {
            ans = append(ans, i)
			count1[i]--
			count2[i]--
        }
    }
    return ans
}
// @lc code=end

