# 2018/11/14 #3
class Solution:
	def groupAnagrams(self, strs):
		if not strs: return []
		res,dic=[],{}
		for s in strs:
			ss=''.join(sorted(s))
			if ss not in dic:
				dic[ss]=[s]
			else:
				dic[ss]+=[s]
		for key in dic:
			# print(dic[key])
			res+=[dic[key]]
		return res

	def isAnagrams(self,strA,strB):
		if len(strA)!=len(strB):
			return False
		# if sorted(strA)==sorted(strB):
		# 	return True
		# else:
		# 	return False
		for s in strB:
			if s not in strA:
				return False
		return True

slt=Solution()
inp=["eat", "tea", "tan", "ate", "nat", "bat"]
ans=[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
print(slt.groupAnagrams(inp))

# print(slt.isAnagrams("eat","tea"))