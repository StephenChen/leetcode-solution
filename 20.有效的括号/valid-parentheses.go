/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (41.94%)
 * Likes:    1673
 * Dislikes: 0
 * Total Accepted:    321.1K
 * Total Submissions: 765.7K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

package leetcode

// @lc code=start
func isValid(s string) bool {
	// region stack

	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	lps := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		p := s[i]
		if _, ok := m[p]; ok {
			lps = append(lps, p)
		} else {
			if len(lps) > 0 && m[lps[len(lps)-1]] == p {
				lps = lps[:len(lps)-1]
			} else {
				return false
			}
		}
	}

	if len(lps) == 0 {
		return true
	}
	return false

	// endregion

	// region other

	// for strings.Contains(s, "{}") || strings.Contains(s, "[]") || strings.Contains(s, "()") {
	// 	s = strings.ReplaceAll(s, "{}", "")
	// 	s = strings.ReplaceAll(s, "[]", "")
	// 	s = strings.ReplaceAll(s, "()", "")
	// }
	// return s == ""

	// endregion
}

// @lc code=end
