# 2018/11/6 #2
class Solution:
	def findDuplicate(self, nums):
		dic={}
		for n in nums:
			if n in dic:
				return n
			else:
				dic[n]=1


slt=Solution()
inp=[1,3,4,2,2]
ans=2
print(slt.findDuplicate(inp)==ans)
inp=[3,1,3,4,2]
ans=3
print(slt.findDuplicate(inp)==ans)
