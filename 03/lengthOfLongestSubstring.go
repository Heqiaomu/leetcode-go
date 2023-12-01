package _3

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	hashTable := map[rune]int{}

	for i, char := range s {
		if index, ok := hashTable[char]; ok && index >= start {
			start = index + 1
		}

		hashTable[char] = i
		length := i - start + 1

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
