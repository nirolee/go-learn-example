package main

func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := last[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last[ch] = i
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	lengthOfLongestSubstring(s)
}
