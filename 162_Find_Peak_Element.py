# 2018/11/6 #3
class Solution:
	def findPeakElement(self, nums):
		if len(nums)==1:
			return 0
		if nums[0]>nums[1]:
			return 0
		for i in range(1,len(nums)-1):
			if nums[i-1]<nums[i]:
				if nums[i]>nums[i+1]:
					return i
		if nums[-1]>nums[-2]:
			return len(nums)-1
		return nums.index(max(nums))

slt=Solution()
inp=[1,2,3,1]
ans=2
print(slt.findPeakElement(inp)==ans)
inp=[1,2,1,3,5,6,4]
ans=1
print(slt.findPeakElement(inp)==ans)
