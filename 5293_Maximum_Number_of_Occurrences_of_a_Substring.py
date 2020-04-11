# 2019/12/22
class Solution:
	def maxFreq(self, s: str, maxLetters: int, minSize: int, maxSize: int) -> int:
		ans=0
		return ans

slt=Solution()
s = "aababcaab"
maxLetters = 2
minSize = 3
maxSize = 4
a=2
print(slt.maxFreq(s,maxLetters,minSize,maxSize)==a)
s = "aaaa"
maxLetters = 1
minSize = 3
maxSize = 3
a=2
print(slt.maxFreq(s,maxLetters,minSize,maxSize)==a)
s = "aabcabcab"
maxLetters = 2
minSize = 2
maxSize = 3
a=3
print(slt.maxFreq(s,maxLetters,minSize,maxSize)==a)
s = "abcde"
maxLetters = 2
minSize = 3
maxSize = 3
a=0
print(slt.maxFreq(s,maxLetters,minSize,maxSize)==a)
