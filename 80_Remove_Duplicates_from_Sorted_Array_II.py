class Solution:
	def removeDuplicates(self, nums):
		i=0
		while i+2<len(nums):
			if nums[i]!=nums[i+1]:
				i+=1
			else:
				if nums[i+1]==nums[i+2]:
					nums.remove(nums[i])
				else:
					i+=1
		return len(nums)

slt=Solution()
nums = [1,1,1,2,2,3]
ans=5
print(slt.removeDuplicates(nums)==ans)
nums = [0,0,1,1,1,1,2,3,3]
ans=7
print(slt.removeDuplicates(nums)==ans)