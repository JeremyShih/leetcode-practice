# 2019/10/25
class Solution:
	def partition(self, s: str):
		ans=[]
		def dfs(l,s):
			if not s:
				ans.append(l)
			else:
				for i in range(1,len(s)+1):
					if s[:i]==s[:i][::-1]:
						dfs(l+[s[:i]],s[i:])
		dfs([],s)
		return ans


slt=Solution()
s="aab"
a=[
  ["aa","b"],
  ["a","a","b"]
]
a.sort()
print(slt.partition(s)==a)
s="aaabb"
a=[["a","a","a","b","b"],["a","a","a","bb"],["a","aa","b","b"],["a","aa","bb"],["aa","a","b","b"],["aa","a","bb"],["aaa","b","b"],["aaa","bb"]]
print(slt.partition(s)==a)
s="aaabbc"
a=[["a","a","a","b","b","c"],["a","a","a","bb","c"],["a","aa","b","b","c"],["a","aa","bb","c"],["aa","a","b","b","c"],["aa","a","bb","c"],["aaa","b","b","c"],["aaa","bb","c"]]
print(slt.partition(s)==a)
s="efe"
a=[["e","f","e"],["efe"]]
# print(slt.partition(s))
print(slt.partition(s)==a)
s="ababbbabbaba"
print(slt.partition(s))
