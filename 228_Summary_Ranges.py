class Solution:
	def summaryRanges(self, nums):
		if len(nums)==0: return []
		i=0
		while i<len(nums)-1 and nums[i]+1==nums[i+1]:
			i+=1
		if i==0:
			return [str(nums[0])]+self.summaryRanges(nums[i+1:])
		else:
			return [str(nums[0])+'->'+str(nums[i])]+self.summaryRanges(nums[i+1:])

slt=Solution()
inp=[0,1,2,4,5,7]
ans=["0->2","4->5","7"]
print(slt.summaryRanges(inp))
inp=[0,2,3,4,6,8,9]
ans=["0","2->4","6","8->9"]
print(slt.summaryRanges(inp))
