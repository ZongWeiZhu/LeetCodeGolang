package array

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		}
		if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, string(cur))
			ptr++
		}
		if cur == ']' {
			sub := []string{}
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-1-i] = sub[len(sub)-1-i], sub[i]
			}
			stack = stack[:len(stack)-1]
			repTime, _ := strconv.Atoi(string(stack[len(stack) - 1]))
			stack = stack[:len(stack)-1]
			t := strings.Repeat(getString(sub), repTime)
			stack = append(stack, t)
			ptr++
		}
	}
	return getString(stack)
}

func getDigits(s string, ptr *int) string {
	digits := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		digits += string(s[*ptr])
	}
	return digits
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
