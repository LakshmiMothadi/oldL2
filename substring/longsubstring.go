package main
func lenOfNoRepeatingSubStr(s string) int {
lastOccurred := make(map[byte]int)
start := 0
maxLen := 0
for i, ch := range []byte(s) {

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
