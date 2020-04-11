# 2019/10/13 contest not solved
class Solution:
	def maxEqualFreq(self, nums: List[int]) -> int:
		dic={}
		for n in nums:
			if n in dic:
				dic[n]+=1
			else:
				dic[n]=0
		for i in range(len(nums)-1, -1, -1):
			dic[nums[i]]-=1

			return i
		return 0

slt=Solution()
nums = [2,2,1,1,5,3,3,5]
ans=7
print(slt.maxEqualFreq(nums)==ans)
nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
ans=13
print(slt.maxEqualFreq(nums)==ans)
nums = [1,1,1,2,2,2]
ans=5
print(slt.maxEqualFreq(nums)==ans)
nums = [10,2,8,9,3,8,1,5,2,3,7,6]
ans=8
print(slt.maxEqualFreq(nums)==ans)
