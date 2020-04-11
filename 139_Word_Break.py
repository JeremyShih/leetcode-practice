# 2019/9/22
class Solution:
	def wordBreak(self, s: str, wordDict: list) -> bool:
		dp=[False]*(len(s)+1)
		dp[0]=True
		for i in range(1,len(s)+1):
			for word in wordDict:
				l=len(word)
				if i+1>=l and word==s[i-l:i]:
					# print(i,word)
					dp[i]=dp[i-l] or dp[i]
		# print(dp)
		return dp[-1]

slt=Solution()
s="leetcode"
wordDict = ["leet", "code"]
print(slt.wordBreak(s,wordDict))
s="applepenapple"
wordDict = ["apple", "pen"]
print(slt.wordBreak(s,wordDict))
s="catsandog"
wordDict = ["cats", "dog", "sando", "and", "cat", "og"]
print(slt.wordBreak(s,wordDict))
s="dogs"
wordDict=["dog","s","gs"]
print(slt.wordBreak(s,wordDict))
