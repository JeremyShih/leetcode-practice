# 2019/8/17
class Solution:
	def lengthOfLongestSubstring(self, s: str) -> int:
		i = 0
		count = 0
		window = ""
		for c in s:
			if c in window:
				window = window[window.index(c) + 1:]
			window += c
			count = max(count, len(window))
		return count

slt = Solution()
s = "abcabcbb"
print(slt.lengthOfLongestSubstring(s)==3)
s = "bbbbb"
print(slt.lengthOfLongestSubstring(s)==1)
s = "pwwkew"
print(slt.lengthOfLongestSubstring(s)==3)
