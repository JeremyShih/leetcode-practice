# 2018/11/13 #1
class Solution:
	def containsNearbyDuplicate(self, nums, k):
		if not nums: return False
		dic={}
		for i,n in enumerate(nums):
			if n in dic and i-dic[n]<=k:
				return True
			else:
				dic[n]=i
		return False


slt=Solution()
inp=[1,2,3,1]
t=3
ans=True
print(slt.containsNearbyDuplicate(inp,t)==ans)
inp=[1,0,1,1]
t=1
ans=True
print(slt.containsNearbyDuplicate(inp,t)==ans)
inp=[1,2,3,1,2,3]
t=2
ans=False
print(slt.containsNearbyDuplicate(inp,t)==ans)
