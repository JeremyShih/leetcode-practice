class Solution:
	def maxProduct(self, nums):
		if not nums:
			return 0
		if len(nums)==1:
			return nums[0]
		max_,min_=1,1
		ans=nums[0]
		for n in nums:
			val=[max_*n,min_*n,1*n]
			max_,min_=max(val),min(val)
			ans=max(max_,ans)
		return ans

slt=Solution()
inp=[2,3,-2,4]
print(slt.maxProduct(inp)==6)
slt=Solution()
inp=[-2,0,-1]
print(slt.maxProduct(inp)==0)
