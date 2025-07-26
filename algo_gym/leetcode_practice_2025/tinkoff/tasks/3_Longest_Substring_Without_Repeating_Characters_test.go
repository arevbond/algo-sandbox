package tasks

import "testing"

func BenchmarkLengthOfLongestSubstringN2(b *testing.B) {
	input := "abcabcbbабвгдежзийклмнопрстуфхцчшщъыьэюя"
	for b.Loop() {
		_ = lengthOfLongestSubstringN2(input)
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	input := "abcabcbbабвгдежзийклмнопрстуфхцчшщъыьэюя"
	for b.Loop() {
		_ = lengthOfLongestSubstring(input)
	}
}

func BenchmarkLengthOfLongestSubstingPAST(b *testing.B) {
	input := "abcabcbbабвгдежзийклмнопрстуфхцчшщъыьэюя"
	for b.Loop() {
		_ = lengthOfLongestSubstringPAST(input)
	}
}
