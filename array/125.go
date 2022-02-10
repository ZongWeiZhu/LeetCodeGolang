package array

import "strings"


/*
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串
*/

func isPalindrome(s string) bool {
	tmpS := ""
	for i := 0; i < len(s); i++ {
		if isNotSpace(s[i]) {
			tmpS += string(s[i])
		}
	}

	tmpS = strings.ToLower(tmpS)
	for i := 0; i < len(tmpS)/2; i++ {
		if tmpS[i] != tmpS[len(tmpS)-i-1] {
			return false
		}
	}

	return true
}

func isNotSpace(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
