package main

import "fmt"

func lenOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[string]string)
	start := 0
	maxLen := 0
	for i, ch := range s {
		if lastIdx, ok := lastOccurred[ch]; ok && lastIdx >= start {
			start = lastIdx + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(lenOfNoRepeatingSubStr("bbbbbbb"))
	fmt.Println(lenOfNoRepeatingSubStr("pwwkew"))
	fmt.Println(lenOfNoRepeatingSubStr(""))
	fmt.Println(lenOfNoRepeatingSubStr("b"))
	fmt.Println(lenOfNoRepeatingSubStr("abcdef"))

}
