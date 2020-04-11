# 2019/11/1
class Solution:
	def singleNumber(self, nums) -> int:
		ans=nums[0]
		for i in range(1,len(nums)):
			ans^=nums[i]
		return ans

slt=Solution()
n=[4,1,2,1,2]
a=4
print(slt.singleNumber(n)==a)
n=[2,2,1]
a=1
print(slt.singleNumber(n)==a)
