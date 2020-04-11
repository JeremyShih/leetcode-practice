# 2018/11/11 #9
class Solution:
	def missingNumber(self, nums):
		l=len(nums)
		total=int(l*(l+1)/2)
		return total-sum(nums)

slt=Solution()
inp=[3,0,1]
ans=2
print(slt.missingNumber(inp)==ans)
inp=[9,6,4,2,3,5,7,0,1]
ans=8
print(slt.missingNumber(inp)==ans)
