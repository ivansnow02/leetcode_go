#include <unordered_map>
#include <vector>
using namespace std;
/*
 * @lc app=leetcode.cn id=454 lang=cpp
 *
 * [454] 四数相加 II
 *
 * https://leetcode.cn/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (64.46%)
 * Likes:    1000
 * Dislikes: 0
 * Total Accepted:    271.9K
 * Total Submissions: 421.8K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n
 * ，请你计算有多少个元组 (i, j, k, l) 能满足：
 *
 *
 * 0 <= i, j, k, l < n
 * nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
 * 输出：2
 * 解释：
 * 两个元组如下：
 * 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) +
 * (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) +
 * (-1) + 0 = 0
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length
 * n == nums2.length
 * n == nums3.length
 * n == nums4.length
 * 1 <= n <= 200
 * -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28
 *
 *
 */

// @lc code=start
class Solution
{
public:
    int
    fourSumCount(vector<int> &nums1, vector<int> &nums2, vector<int> &nums3,
                 vector<int> &nums4)
    {
        unordered_map<int, int> umap;
        auto count = 0;
        for (auto a : nums1)
        {
            for (auto b : nums2)
            {
                umap[a + b]++;
            }
        }
        for (auto c : nums3)
        {
            for (auto d : nums4)
            {
                if (umap.find(0 - (c + d)) != umap.end()) 
                {
                    count += umap[0 - (c + d)];
                }
            }
        }
        return count;
    }
};
// @lc code=end
